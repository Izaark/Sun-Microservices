// Package config loads/defines main configurations to run this microservice
package config

import (
	"errors"

	"github.com/joho/godotenv"
)

// FunInitConfig load environment variables
func FunInitConfig() error {
	err := godotenv.Load("environment.env")
	if err != nil {
		err = errors.New("FunInitConfig: couldn't initialize environment -> " + err.Error())
		return err
	}
	return nil
}
