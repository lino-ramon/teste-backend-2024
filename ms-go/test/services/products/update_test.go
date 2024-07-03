package products

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"ms-go/config/logger"
	"ms-go/db"
)


func TestUpdate(t *testing.T) {
	const (
		LOG_LEVEL = "debug"
		LOG_PATH  = "../../../logs/update_test.log"
	)
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()

	// Caso de teste 1: Atualização válida
	t.Run("Update Valid Product", func(t *testing.T) {
		// Preparação do banco de dados com um produto inicial
		initialProduct := models.Product{
			ID:          1,
			Name:        "Initial Product",
			Brand:       "Initial Brand",
			Price:       50.0,
			Description: "Initial Description",
			Stock:       20,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err := db.Connection().InsertOne(context.Background(), initialProduct)
		assert.NoError(t, err, "Error inserting initial product")

		updatedProduct := models.Product{
			ID:          initialProduct.ID,
			Name:        "Updated Product",
			Brand:       "Updated Brand",
			Price:       75.0,
			Description: "Updated Description",
			Stock:       30,
			CreatedAt:   initialProduct.CreatedAt,
			UpdatedAt:   time.Now(),
		}

		updatedResult, err := products.Update(updatedProduct, false)

		assert.NoError(t, err, "Unexpected error when updating product")
		assert.NotNil(t, updatedResult, "Product update returned nil")
		assert.Equal(t, updatedProduct.Name, updatedResult.Name, "Updated product name does not match")

		var result models.Product
		err = db.Connection().FindOne(context.Background(), bson.M{"id": updatedProduct.ID}).Decode(&result)
		assert.NoError(t, err, "Error finding updated product in database")
		assert.Equal(t, updatedProduct.Name, result.Name, "Recovered product name does not match")
	})

	// Caso de teste 2: ID do produto não especificado
	t.Run("Missing Product ID", func(t *testing.T) {
		invalidProduct := models.Product{
			Name:        "Invalid Product",
			Brand:       "Invalid Brand",
			Price:       100.0,
			Description: "Invalid Description",
			Stock:       15,
		}

		_, err := products.Update(invalidProduct, false)

		assert.Error(t, err, "Expected an error when updating with missing ID")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Returned error is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})

	// Caso de teste 3: Produto não encontrado
	t.Run("Product Not Found", func(t *testing.T) {
		notFoundProduct := models.Product{
			ID:          999, // Supondo um ID que não existe no banco de dados
			Name:        "Product Not Found",
			Brand:       "Brand Not Found",
			Price:       200.0,
			Description: "Description Not Found",
			Stock:       5,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err := products.Update(notFoundProduct, false)

		assert.Error(t, err, "Expected an error when updating non-existent product")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Returned error is not of type GenericError")
		assert.Equal(t, http.StatusNotFound, genericErr.Code, "Error code does not match")
	})

	// Caso de teste 4: Atualização inválida - Preço negativo
	t.Run("Invalid Update - Negative Price", func(t *testing.T) {
		invalidPriceProduct := models.Product{
			ID:          1, // ID de um produto existente para teste
			Name:        "Invalid Price Product",
			Brand:       "Invalid Price Brand",
			Price:       -10.0,
			Description: "Invalid Price Description",
			Stock:       25,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err := products.Update(invalidPriceProduct, false)

		assert.Error(t, err, "Expected an error when updating with negative price")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Returned error is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})
}
