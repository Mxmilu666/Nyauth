package main

import (
	"nyauth_backed/source"
	"nyauth_backed/source/logger"
	"nyauth_backed/source/server"
)

func main() {
	logger.InitLogger(logger.DEBUG)
	logger.Info("Hello, Nyauth!")
	source.LoadConfig()
	server.Setupserver()
}
