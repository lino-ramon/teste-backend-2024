package main

import (
	"ms-go/app/consumers"
	"ms-go/config/logger"
	_ "ms-go/db"
	"ms-go/router"
)

const (
	LOG_LEVEL = "debug"
	LOG_PATH  = "../../logs/product_test.log"
)

func main() {
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()
	// Iniciando Consumers
	logger.Info("Main - Starting Consumers")
	go consumers.StartConsumers()

	logger.Info("Main - Starting Routers")
	// Iniciando API
	router.Run()
}
