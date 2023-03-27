package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestNewProduct (t *testing.T) {
	product, err := NewProduct("Product 1", 1200)

	assert.Nil(t,err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 1200.00, product.Price)
}

func TestProductWhenNameIsRequired (t *testing.T) {
	product, err := NewProduct("", 1200)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired (t *testing.T) {
	product, err := NewProduct("Product 2", 0)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid (t *testing.T) {
	product, err := NewProduct("Product 2", -10)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product 1", 1200)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}


