package consumers

import (
	"ms-go/config/logger"
)

func StartConsumers() {
	logger.Info("Consumers - Starting Consumers")
	StartRailsToGoConsumer()
	// add other consumers..
}
