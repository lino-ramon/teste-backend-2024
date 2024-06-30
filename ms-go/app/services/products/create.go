package products

import (
	"context"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"
	"time"

	"ms-go/config/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(data models.Product, isAPI bool) (*models.Product, error) {
	logger.Info("ProductsService - Creating Product")
	if data.ID == 0 {
		var max models.Product

		opts := options.FindOne()

		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

		db.Connection().FindOne(context.TODO(), bson.D{}, opts).Decode(&max)

		data.ID = max.ID + 1
	}

	if err := data.Validate(); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusUnprocessableEntity}
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = data.CreatedAt

	logger.Info("ProductsService - Saving Product")
	_, err := db.Connection().InsertOne(context.TODO(), data)

	if err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	logger.Info("ProductsService - Product Saved sucessfull")
	defer db.Disconnect()

	if isAPI {
	}

	return &data, nil
}
