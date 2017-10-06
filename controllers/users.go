package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/sun/Sun-Microservices/models"
)

func UsersRouter() {

	// Execution mode. Exotic condition to ensure ONLY DEBUG/RELEASE values are accepted
	if os.Getenv("SUN_ENV_DEPLOY_MODE") == "DEBUG" || os.Getenv("SUN_ENV_DEPLOY_MODE") == "RELEASE" {
		if os.Getenv("SUN_ENV_DEPLOY_MODE") == "RELEASE" {
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		fmt.Println("*ERROR Posts Microservice Router: invalid environment 'DEPLOY_MODE' aborting!")
		return
	}
	// Router config
	router := gin.Default()

	// CORS config
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE ,OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// Routes/Endpoints
	user := router.Group("api")
	{
		user.GET("/users", handlerGetAllUsers)
		user.POST("/user/register", handlerPostUser)
	}

	router.Run(":" + os.Getenv("SUN_ENV_API_PORT"))
}

func handlerPostUser(c *gin.Context) {
	var vginResponse gin.H
	var vUser models.ObjUserPost
	var err error

	// Binding provided payload with corresponding struct
	err = c.BindJSON(&vUser)
	if err != nil {
		err = errors.New("*ERROR handlerPostGroup: couldn't bind payload provided with Group struct -> " + err.Error())
		fmt.Println(err)
		vginResponse = gin.H{"message": "error reading payload provided", "response": nil, "error": "RE", "status": http.StatusBadRequest}
		c.JSON(http.StatusBadRequest, vginResponse)
		return
	}
	err = models.FunPostUser(vUser)
	if err != nil {
		err = errors.New("*ERROR handlerPostUser: couldn't register user -> " + err.Error())
		fmt.Println(err)
		vginResponse = gin.H{"message": "internal error", "response": nil, "error": "IE", "status": http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, vginResponse)
		return
	}
	vginResponse = gin.H{"message": "user successfully registered", "error": nil, "status": http.StatusOK}
	c.JSON(http.StatusOK, vginResponse)
}
func handlerGetAllUsers(c *gin.Context) {
	var response gin.H
	var users []models.ObjUserGet
	var err error

	users, err = models.FunGetAllUser()
	for user, data := range users {
		fmt.Println(data.Name)
		users[user] = data

	}

	if err != nil {
		response = gin.H{"message": "error quering users", "response": users}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if users == nil {
		response = gin.H{"message": "not users found", "response": users}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response = gin.H{"message": "users found", "response": users}
	c.JSON(http.StatusOK, response)

}
