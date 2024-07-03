package products

import (
	"ms-go/config/logger"
	"context"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"
	"testing"
	"time"

	"ms-go/app/services/products"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)


func TestCreate(t *testing.T) {
	const (
		LOG_LEVEL = "debug"
		LOG_PATH  = "../../../logs/create_test.log"
	)
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()
	
	// Caso de teste 1: Produto novo
	t.Run("Create New Product", func(t *testing.T) {
		err := db.Connection().Drop(context.Background())
		assert.NoError(t, err, "Error clearing collection")

		product := models.Product{
			Name:        "Test Product",
			Brand:       "Test Brand",
			Price:       99.99,
			Description: "Test Description",
			Stock:       10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		createdProduct, err := products.Create(product, false)

		assert.NoError(t, err, "Unexpected error when creating the product")
		assert.NotNil(t, createdProduct, "Product not created correctly")
		assert.Equal(t, product.Name, createdProduct.Name, "Product name does not match")

		var result models.Product
		err = db.Connection().FindOne(context.Background(), bson.M{"name": product.Name}).Decode(&result)
		assert.NoError(t, err, "Error finding product in database")
		assert.Equal(t, product.Name, result.Name, "Recovered product name does not match")
	})

	// Caso de teste 2: Erro na validação do produto - Produto sem nome
	t.Run("Validation Error", func(t *testing.T) {
		// Dados do produto inválidos (sem nome)
		product := models.Product{
			Brand:       "Test Brand",
			Price:       99.99,
			Description: "Test Description",
			Stock:       10,
		}

		_, err := products.Create(product, false)

		assert.Error(t, err, "A validation error was expected")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Error returned is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})

	// Caso de teste 3: Erro na validação do produto - Preco invalido
	t.Run("Validation Error", func(t *testing.T) {
		product := models.Product{
			Name: 		 "Test Product",	
			Brand:       "Test Brand",
			Price:       -10.99,
			Description: "Test Description",
			Stock:       10,
		}

		_, err := products.Create(product, false)

		assert.Error(t, err, "A validation error was expected")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Error returned is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})

	// Caso de teste 3: Erro na validação do produto - Stock invalido
	t.Run("Validation Error", func(t *testing.T) {
		product := models.Product{
			Name: 		 "Test Product",	
			Brand:       "Test Brand",
			Price:       99.99,
			Description: "Test Description",
			Stock:       -1,
		}

		_, err := products.Create(product, false)

		assert.Error(t, err, "A validation error was expected")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "Error returned is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})
}
