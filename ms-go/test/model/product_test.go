package model

import (
	"ms-go/config/logger"
	"ms-go/app/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	LOG_LEVEL = "debug"
	LOG_PATH  = "../../logs/product_test.log"
)

func TestProductValidation(t *testing.T) {
	logger.InitLogger(LOG_LEVEL, LOG_PATH)
	defer logger.Logger.Sync()

	validProduct := models.Product{
		ID:          1,
		Name:        "Produto Teste",
		Brand:       "Marca Teste",
		Price:       99.99,
		Description: "Descrição Teste",
		Stock:       10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := validProduct.Validate()
	assert.Nil(t, err, "Expected no validation error for a valid product")

	// Cenário: Nome do produto com menos de 4 caracteres
	invalidProductName := validProduct
	invalidProductName.Name = "Abc"
	err = invalidProductName.Validate()
	assert.NotNil(t, err, "Expected validation error for product with short name")

	// Cenário: Preço do produto menor que 0.01
	invalidProductPrice := validProduct
	invalidProductPrice.Price = 0.001
	err = invalidProductPrice.Validate()
	assert.NotNil(t, err, "Expected validation error for product with low price")

	// Cenário: Estoque do produto menor que 0
	invalidProductStock := validProduct
	invalidProductStock.Stock = -1
	err = invalidProductStock.Validate()
	assert.NotNil(t, err, "Expected validation error for product with negative stock")

	// Cenário: Marca do produto vazia
	invalidProductBrand := validProduct
	invalidProductBrand.Brand = ""
	err = invalidProductBrand.Validate()
	assert.NotNil(t, err, "Expected validation error for product with empty brand")

	// Cenário: Descrição do produto vazia
	invalidProductDescription := validProduct
	invalidProductDescription.Description = ""
	err = invalidProductDescription.Validate()
	assert.NotNil(t, err, "Expected validation error for product with empty description")
}
