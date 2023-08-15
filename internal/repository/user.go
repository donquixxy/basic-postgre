package repository

import (
	"postgre-basic/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, value *domain.User) (*domain.User, error)
	FindByName(db *gorm.DB, name string) (*domain.User, error)
	FindAll(db *gorm.DB) []*domain.User
	Update(db *gorm.DB, newValue *domain.User) error
	FindByID(db *gorm.DB, id string) (*domain.User, error)
	UpdateZeroAge(db *gorm.DB, newValue *domain.User) error
	Delete(db *gorm.DB, valueDelete *domain.User) error
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

func (*UserRepositoryImpl) FindAll(db *gorm.DB) []*domain.User {
	var ent []*domain.User

	db.Preload("Company").
		Find(&ent)

	return ent
}

func (*UserRepositoryImpl) UpdateZeroAge(db *gorm.DB, newValue *domain.User) error {
	res := db.Where("id = ?", newValue.ID).Select("age").Updates(newValue)

	return res.Error
}

func (*UserRepositoryImpl) Update(db *gorm.DB, newValue *domain.User) error {
	res := db.Where("id = ?", newValue.ID).Select("age").Updates(newValue)
	return res.Error
}

func (*UserRepositoryImpl) FindByID(db *gorm.DB, id string) (*domain.User, error) {
	user := &domain.User{}

	res := db.Where("id = ?", id).Preload("Company").
		First(&user)

	return user, res.Error
}

func (*UserRepositoryImpl) Delete(db *gorm.DB, valueDelete *domain.User) error {
	res := db.Delete(&valueDelete)

	return res.Error
}
