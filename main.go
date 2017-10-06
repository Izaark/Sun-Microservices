// Package main initializes general configurations and controllers to API's access
package main

import (
	"fmt"
	"log"

	"github.com/sun/Sun-Microservices/config"
	"github.com/sun/Sun-Microservices/controllers"
)

func init() {
	err := config.FunInitConfig()
	if err != nil {
		log.Fatal("*ERROR init: couldn't initialize configuration -> ", err.Error())
	}

	vsessionSocial, err := config.FunOpenDatabaseConnection("social")
	defer vsessionSocial.Close()
	if err != nil {
		fmt.Println("*ERROR init: couldn't connect database -> ", err.Error())
	}

	vsessionMetrics, err := config.FunOpenDatabaseConnection("metricas")
	defer vsessionMetrics.Close()
	if err != nil {
		fmt.Println("*ERROR init: couldn't connect database -> ", err.Error())
	}
}

func main() {
	controllers.UsersRouter()
}
