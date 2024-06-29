package consumers

import (
	"ms-go/app/consumers/kafka"
	"ms-go/config/logger"
)

func StartConsumers() {
	logger.Info("Consumers - Starting Consumers")
	kafka.StartRailsToGoConsumer()
	// add other consumers..
}