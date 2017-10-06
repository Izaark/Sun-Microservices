package models

import (
	"errors"
	"github.com/sun/Sun-Microservices/config"
	r "gopkg.in/gorethink/gorethink.v3"
)

// Constants
const CstrNinTable = "sun_nin"
const CstrUserTable = "users"

type ObjUserPost struct {
	Name     string `json:"name,omitempty" binding:"required" gorethink:"name"`
	LastName string `json:"lastname" gorethink:"lastname"`
	Age      int    `json:"age" gorethink:"age"`
}
type ObjUserGet struct {
	Name     string `json:"name,omitempty" binding:"required" gorethink:"name"`
	LastName string `json:"lastname" gorethink:"lastname"`
	Age      int    `json:"age" gorethink:"age"`
}

//Post a user whit information objUserPost and the json from payload
func FunPostUser(vUser ObjUserPost) error {

	// Open database connection
	vsessionSocial, err := config.FunOpenDatabaseConnection(CstrNinTable)
	defer vsessionSocial.Close()
	if err != nil {
		err = errors.New("*ERROR FunPostGroup: couldn't connect database -> " + err.Error())
		return err
	}

	_, err = r.Table(CstrUserTable).Insert(vUser).RunWrite(vsessionSocial)
	if err != nil {
		err = errors.New("*ERROR FunPostGroup: could'n insert new group " + vUser.Name + " -> " + err.Error())
		return err
	}

	return nil
}

//Get all users of table users
func FunGetAllUser() ([]ObjUserGet, error) {

	var users []ObjUserGet
	var cursor *r.Cursor

	vsessionSocial, err := config.FunOpenDatabaseConnection(CstrNinTable)
	defer vsessionSocial.Close()
	if err != nil {
		err = errors.New("*ERROR FunGetAllusers: can't connect database" + " --> " + err.Error())
		return users, err
	}

	cursor, err = r.Table(CstrUserTable).Run(vsessionSocial)
	defer cursor.Close()
	if err != nil {
		err = errors.New("*ERROR FunGetAllusers: can't retrieve users" + " --> " + err.Error())
		return users, err
	}

	err = cursor.All(&users)
	if err != nil {
		err = errors.New("*ERROR FunGetAllusers: can't use cursor to retrieve users" + " --> " + err.Error())
		return users, err
	}

	return users, nil

}
