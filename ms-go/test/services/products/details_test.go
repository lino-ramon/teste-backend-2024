package products

import (
	"context"
	"net/http"
	"testing"
	"ms-go/config/logger"

	"github.com/stretchr/testify/assert"

	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"ms-go/db"
)

func TestDetails(t *testing.T) {
	const (
		LOG_LEVEL = "debug"
		LOG_PATH  = "../../../logs/details_test.log"
	)
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()
	// Caso de teste 1: Detalhes de um produto existente
	t.Run("Existing Product Details", func(t *testing.T) {
		initialProduct := models.Product{
			ID:          1,
			Name:        "Test Product",
			Brand:       "Test Brand",
			Price:       99.99,
			Description: "Test Description",
			Stock:       10,
		}

		_, err := db.Connection().InsertOne(context.Background(), initialProduct)
		assert.NoError(t, err, "Error inserting initial product")

		result, err := products.Details(models.Product{ID: initialProduct.ID})

		assert.NoError(t, err, "Unexpected error getting product details")
		assert.NotNil(t, result, "Product details must not be null")
		assert.Equal(t, initialProduct.Name, result.Name, "The name of the retrieved product does not match")
	})

	// Caso de teste 2: Produto não encontrado
	t.Run("Product Not Found", func(t *testing.T) {
		_, err := products.Details(models.Product{ID: 999})

		assert.Error(t, err, "An error was expected when searching for a non-existent product")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "The error returned is not of type GenericError")
		assert.Equal(t, http.StatusNotFound, genericErr.Code, "Error code does not match")
	})

	// Caso de teste 3: Parâmetros inválidos (ID não especificado)
	t.Run("Invalid Parameters", func(t *testing.T) {
		_, err := products.Details(models.Product{})

		assert.Error(t, err, "An invalid parameters error was expected")
		genericErr, ok := err.(*helpers.GenericError)
		assert.True(t, ok, "The error returned is not of type GenericError")
		assert.Equal(t, http.StatusUnprocessableEntity, genericErr.Code, "Error code does not match")
	})
}
