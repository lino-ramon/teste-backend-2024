package consumers

import (
	"context"
	"encoding/json"
	"time"

	"ms-go/app/models"
	"ms-go/app/services/products"
	"ms-go/config/logger"

	"github.com/segmentio/kafka-go"
)

const (
	TOPIC_GROUP = "ms-go-group"
	TOPIC_NAME  = "rails-to-go"
	MIN_BYTES   = 10e3 // 10KB
	MAX_BYTES   = 10e6 // 10MB
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
			logger.Error("RailsToGo - Error fetch message: ", err)
			time.Sleep(RETRY_DELAY)
			continue
		}

		var product models.Product
		err = json.Unmarshal(message.Value, &product)
		if err != nil {
			logger.Error("RailsToGo - Error unmarshalling message: ", err)
			continue
		}

		logger.Debug("RailsToGo - Received Rails Product:", product)
		err = saveProductRails(product)
		if err != nil {
			logger.Error("RailsToGo - Error to Save Product: ", err)
			continue
		}
		commitMessages(reader, message)
	}
}

func commitMessages(reader *kafka.Reader, message kafka.Message) {
	logger.Info("RailsToGo - Commiting Messages")
	if err := reader.CommitMessages(context.Background(), message); err != nil {
		logger.Error("RailsToGo - Failed to commit message:", err)
	}
}

func saveProductRails(product models.Product) error {
	_, err := products.Details(product)
	if err == nil {
		_, err = products.Update(product, false)
	} else {
		_, err = products.Create(product, false)
	}
	return err
}
