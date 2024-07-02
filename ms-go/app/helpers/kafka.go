package helpers

import (
	"encoding/json"
	"ms-go/app/producers"
	"ms-go/config/logger"
	"net/http"
)

func ProduceProductMessage(data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal product:", err)
		return &GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	err = producers.ProduceMessage(payload)
	if err != nil {
		logger.Error("Failed to produce Kafka message:", err)
		return &GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	return nil
}
