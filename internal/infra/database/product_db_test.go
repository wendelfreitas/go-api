package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wendelfreitas/go-api/api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	// Migrate the schema
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)

	assert.NoError(t,err)
	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.NoError(t,err)
	assert.Equal(t, product.Price, 10.0)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) 
	for i:=1; i< 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64() * 100)
		assert.NoError(t,err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1,10,"asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2,10,"asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3,10,"asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)

}


func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) 
	p, err := entity.NewProduct("Product 1", 10)

	assert.NoError(t,err)

	productDB := NewProduct(db)

	err = productDB.Create(p)

	if err != nil {
		t.Error(err)
	}

	product, err := productDB.FindByID(p.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Name, product.Name)
	assert.Equal(t, p.Price, product.Price)
}


func TestUpdateProduct (t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	p, err := entity.NewProduct("Product 1", 10)

	assert.NoError(t, err)

	productDB := NewProduct(db)


	p.Name = "Product Edited"

	err = productDB.Update(p)
	assert.NoError(t,err)

	product, err := productDB.FindByID(p.ID.String())

	assert.NoError(t, err)

	assert.Equal(t, product.Name, "Product Edited")
}

func TestDeleteProduct(t *testing.T) { 
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	p, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t,err)
	db.Create(p)
	productDB := NewProduct(db)

	err = productDB.Delete(p.ID.String())	
	assert.NoError(t, err)

	product, err := productDB.FindByID(p.ID.String())

	assert.Error(t, err)
	assert.Nil(t, product)

}