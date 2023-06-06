package repository

import (
	"postgre-basic/internal/domain"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(db *gorm.DB, value *domain.Company) (*domain.Company, error)
	FindAll(db *gorm.DB) []*domain.Company
}

type CompanyRepositoryImpl struct{}

func NewCompanyRepository() CompanyRepository {
	return &CompanyRepositoryImpl{}
}

func (*CompanyRepositoryImpl) Create(db *gorm.DB, value *domain.Company) (*domain.Company, error) {
	res := db.Create(&value)

	return value, res.Error
}

func (*CompanyRepositoryImpl) FindAll(db *gorm.DB) []*domain.Company {
	var ent []*domain.Company

	db.Find(&ent)

	return ent
}
