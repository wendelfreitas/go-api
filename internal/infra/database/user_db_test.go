package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wendelfreitas/go-api/api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "john@doe.com", "1234567")

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t,err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.NotEmpty(t, userFound.ID)
	assert.NotEmpty(t, userFound.Password)
	assert.Equal(t, "John Doe", userFound.Name)
	assert.Equal(t, "john@doe.com", userFound.Email)
}

func TestFindByEmail (t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john@doe2.com", "1234567")

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t,err)
	
	userFound, err := userDB.FindByEmail(user.Email)
	
	assert.Nil(t,err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

