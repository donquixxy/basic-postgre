package usecases

import (
	"postgre-basic/internal/api/request/companyrequest"
	"postgre-basic/internal/domain"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/repository"
	"postgre-basic/utils/etc"
	"time"

	"gorm.io/gorm"
)

type CompanyServices interface {
	Create(request *companyrequest.CreateRequest) (*domain.Company, error)
}

type CompanyServicesImpl struct {
	CompanyRepository repository.CompanyRepository
	Database          *gorm.DB
}

func NewCompanyServices(
	CompanyRepository repository.CompanyRepository,
	Database *gorm.DB,
) CompanyServices {
	return &CompanyServicesImpl{
		CompanyRepository: CompanyRepository,
		Database:          Database,
	}
}

func (this *CompanyServicesImpl) Create(request *companyrequest.CreateRequest) (*domain.Company, error) {
	companyEnt := &domain.Company{
		ID:        etc.GenerateRandomUUID(),
		Name:      request.Name,
		Phone:     request.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := this.CompanyRepository.Create(this.Database, companyEnt)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return created, nil
}
