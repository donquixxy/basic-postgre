package repository

import (
	"postgre-basic/database/postgre"
	"postgre-basic/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGetByID(t *testing.T) {
	db, err := postgre.InitPostgreDb()
	assert.NoError(t, err)

	database, er := db.DB()

	assert.NoError(t, er)

	defer database.Close()

	expected := &domain.User{
		ID:   "321",
		Name: "indah",
		Age:  21,
	}

	userRepo := UserRepositoryImpl{}
	result, err := userRepo.FindByID(db, "321")
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, result.ID)
}

func TestGetAllUsers(t *testing.T) {
	db, err := postgre.InitPostgreDb()
	assert.NoError(t, err)

	database, er := db.DB()

	assert.NoError(t, er)

	defer database.Close()

	expected := 4

	userRepo := UserRepositoryImpl{}

	listData := userRepo.FindAll(db)

	assert.NotEmpty(t, listData)
	assert.Equal(t, expected, len(listData))
}
