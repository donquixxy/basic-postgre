package repository

import (
	"postgre-basic/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, value *domain.User) (*domain.User, error)
	FindByName(db *gorm.DB, name string) (*domain.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (*UserRepositoryImpl) Create(db *gorm.DB, value *domain.User) (*domain.User, error) {
	err := db.Create(&value)

	return value, err.Error
}

func (*UserRepositoryImpl) FindByName(db *gorm.DB, name string) (*domain.User, error) {
	user := &domain.User{}

	responses := db.Where("name LIKE ?", name).First(&user)

	return user, responses.Error
}
