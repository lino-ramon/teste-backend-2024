package main

import (
	_ "ms-go/db"
	"ms-go/router"
	"ms-go/app/consumers"
	"ms-go/config/logger"
)

func main() {
	logger.InitLogger("debug")
	defer logger.Logger.Sync()
	// Iniciando Consumers
	logger.Info("Main - Starting Consumers")
	go consumers.StartConsumers()

	logger.Info("Main - Starting Routers")
	// Iniciando API
	router.Run()
}
