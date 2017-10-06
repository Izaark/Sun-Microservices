// Package config loads/defines main configurations to run this microservice
package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
)

// FunOpenDatabaseConnection initialize a connection with the provided database, returns the corresponding session object
func FunOpenDatabaseConnection(pstrDatabase string) (*r.Session, error) {
	var vSession *r.Session
	var vstrDatabase string
	var vintTimeoutSecs int
	var err error

	vstrDatabase = os.Getenv("SUN_ENV_NIN_DATABASE")

	// Preventing error while parsing timeout from environment variables
	vintTimeoutSecs, err = strconv.Atoi(os.Getenv("CDMS_ENV_SOCIAL_DATABASE_TIMEOUT_SECS"))
	if err != nil {
		vintTimeoutSecs = 1
	}

	// Creating connection with defined parameters
	vSession, err = r.Connect(r.ConnectOpts{
		Address:  os.Getenv("SUN_ENV_NIN_DATABASE_ADDRESS"),
		Database: vstrDatabase,
		Timeout:  time.Duration(vintTimeoutSecs) * time.Second,
	})
	if err != nil {
		err = errors.New("*ERROR FunOpenDatabaseConnection: couldn't connect to rethinkdb -> " + err.Error())
		return vSession, err
	}

	result, err := r.Expr("-> Database successfully connected").Run(vSession)
	defer result.Close()
	if err != nil {
		err = errors.New("*ERROR FunOpenDatabaseConnection: couldn't connect " + vstrDatabase + " database -> " + err.Error())
		return vSession, err
	}

	response := ""
	err = result.One(&response)
	if err != nil {
		err = errors.New("*ERROR FunOpenDatabaseConnection: " + vstrDatabase + " database is not responding -> " + err.Error())
		return vSession, err
	}

	//fmt.Println(response)
	return vSession, nil
}
