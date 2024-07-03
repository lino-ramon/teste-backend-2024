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
)

func Update(data models.Product, isAPI bool) (*models.Product, error) {
	logger.Info("ProductsService - Updating Product")
	if data.ID == 0 {
		logger.Error("ProductsService - Missing parameter ID")
		return nil, &helpers.GenericError{Msg: "Missing parameters", Code: http.StatusUnprocessableEntity}
	}

	var product models.Product

	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product); err != nil {
		logger.Error("ProductsService - Product Not Found")
		return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
	}

	data.UpdatedAt = time.Now()

	setUpdate(&data, &product)

	if err := data.Validate(); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusUnprocessableEntity}
	}
	
	logger.Info("ProductsService - Updating Product on MongoDB")
	if err := db.Connection().FindOneAndUpdate(context.TODO(), bson.M{"id": data.ID}, bson.M{"$set": data}).Decode(&product); err != nil {
		logger.Error("ProductsService - Error to update product: ", err)
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	logger.Info("ProductsService - Validating Update")
	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product); err != nil {
		logger.Error("ProductsService - Product Not Found")
		return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
	}
	logger.Info("ProductsService - Product updated sucessfull")
	defer db.Disconnect()

	if isAPI {
		err := helpers.ProduceProductMessage(data)
		if err != nil {
			logger.Error("Failed to produce Kafka message:", err)
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
		}
	}

	return &product, nil
}

func setUpdate(new, old *models.Product) {
	if new.ID == 0 {
		new.ID = old.ID
	}

	if new.Name == "" {
		new.Name = old.Name
	}

	if new.Brand == "" {
		new.Brand = old.Brand
	}

	if new.Price == 0 {
		new.Price = old.Price
	}

	if new.Description == "" {
		new.Description = old.Description
	}

	if new.Stock == 0 {
		new.Stock = old.Stock
	}

	new.CreatedAt = old.CreatedAt

	new.UpdatedAt = time.Now()
}
