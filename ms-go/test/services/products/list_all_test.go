package products

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"ms-go/config/logger"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"ms-go/db"
)

func TestListAll(t *testing.T) {
	const (
		LOG_LEVEL = "debug"
		LOG_PATH  = "../../../logs/list_all_test.log"
	)
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()
	// Caso de teste 1: Listagem de produtos existentes
	t.Run("List Existing Products", func(t *testing.T) {
		initialProducts := []models.Product{
			{ID: 1, Name: "Product 1", Description: "Description 1", Price: 10.99, Stock: 5},
			{ID: 2, Name: "Product 2", Description: "Description 1", Price: 19.99, Stock: 8},
		}
		err := db.Connection().Drop(context.Background())
		assert.NoError(t, err, "Error clearing collection")
		
		for _, prod := range initialProducts {
			_, err := db.Connection().InsertOne(context.Background(), prod)
			assert.NoError(t, err, "Error when inserting initial products")
		}

		products, err := products.ListAll()

		assert.NoError(t, err, "Unexpected error when listing products")
		assert.NotNil(t, products, "The product list must not be null")
		assert.Equal(t, len(initialProducts), len(products), "The number of products in the list does not match")

		for i, expected := range initialProducts {
			assert.Equal(t, expected.Name, products[i].Name, "Product name does not match")
			assert.Equal(t, expected.Price, products[i].Price, "The product price does not match")
			assert.Equal(t, expected.Stock, products[i].Stock, "Product stock does not match")
		}
	})

	// Caso de teste 2: Listagem de produtos quando não há produtos
	t.Run("List Empty Products", func(t *testing.T) {
		_, err := db.Connection().DeleteMany(context.Background(), bson.D{})
		assert.NoError(t, err, "Error cleaning products from database")

		products, err := products.ListAll()

		assert.NoError(t, err, "Unexpected error when listing products")
		assert.NotNil(t, products, "The product list must not be null")
		assert.Empty(t, products, "The product list must be empty")
	})
}
