package kafka

import (
	"encoding/json"
	"context"
	"time"

	"github.com/segmentio/kafka-go"
    "ms-go/app/services/products"
	"ms-go/config/logger"
	"ms-go/app/models"
)

const (
	TOPIC_GROUP = "ms-go-group"
	TOPIC_NAME  = "rails-to-go"
	MIN_BYTES   = 10e3 // 10KB
	MAX_BYTES   = 10e6 // 10MB
	RETRY_COUNT = 5
	RETRY_DELAY = 10 * time.Second
)

func createKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:29092"},
		GroupID:  TOPIC_GROUP,
		Topic:    TOPIC_NAME,
		MinBytes: MIN_BYTES,
		MaxBytes: MAX_BYTES,
	})
}

func StartRailsToGoConsumer() {
	reader := createKafkaReader()
	defer reader.Close()

	for {
		logger.Info("RailsToGo - Fetching Messages")
                
        message, err := reader.FetchMessage(context.Background())
        if err != nil {
            logger.Error("RailsToGo - Error fetching message: ", err)
            time.Sleep(RETRY_DELAY)
            continue
        }

		var product models.Product
		err = json.Unmarshal(message.Value, &product)
		if err != nil {
			logger.Error("RailsToGo - Error unmarshalling message: ", err)
			continue
		}
		
        logger.Debug("RailsToGo - Received product:", product)
        _, err = products.Create(product, false)
        if err != nil {
            logger.Error("RailsToGo - Error to Create Product: ", err)
            continue
        }

        commitMessages(reader, message)
	}
}

func commitMessages(reader *kafka.Reader, message kafka.Message) {
    if err := reader.CommitMessages(context.Background(), message); err != nil {
        logger.Error("RailsToGo - Failed to commit message:", err)
    }
}