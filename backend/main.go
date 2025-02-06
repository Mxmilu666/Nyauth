package main

import (
	"fmt"
	"nyauth_backed/source"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/logger"
	"nyauth_backed/source/server"
)

func main() {
	logger.InitLogger(logger.DEBUG)
	logger.Info("Hello, Nyauth!")
	err := source.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	err = database.InitDatabase()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to initialize database: %s\n", err.Error()))
	}
	_, err = helper.GetInstance()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to initialize JWTHelper: %s\n", err.Error()))
	}
	server.Setupserver()
}
