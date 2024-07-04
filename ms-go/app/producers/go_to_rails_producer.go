package producers

import (
	"context"
	"ms-go/config/logger"

	"github.com/segmentio/kafka-go"
)

const (
	KAFKA_BROKER = "kafka:29092"
	TOPIC_NAME   = "go-to-rails"
)

func createKafkaWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{KAFKA_BROKER},
		Topic:   TOPIC_NAME,
	})
}

func ProduceMessage(payload []byte) error {
	logger.Info("Producing Kafka Message")
	writer := createKafkaWriter()
	defer writer.Close()

	msg := kafka.Message{
		Value: payload,
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		logger.Error("Failed to produce message:", err)
		return err
	}

	logger.Info("Message produced to Kafka topic:", TOPIC_NAME)
	return nil
}
