package usecases

import (
	"fmt"
	"postgre-basic/internal/api/request/userrequest"
	"postgre-basic/internal/domain"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/repository"
	"time"

	"gorm.io/gorm"
)

type UserServices interface {
	CreateUser(request *userrequest.CreateRequest) (*domain.User, error)
}

type UserServicesImpl struct {
	Database       *gorm.DB
	UserRepository repository.UserRepository
}

func NewUserServices(
	Database *gorm.DB,
	UserRepository repository.UserRepository,
) UserServices {
	return &UserServicesImpl{
		Database:       Database,
		UserRepository: UserRepository,
	}
}

func (this *UserServicesImpl) CreateUser(request *userrequest.CreateRequest) (*domain.User, error) {

	// Check if user is duplicate
	duplicatedUser, _ := this.UserRepository.FindByName(this.Database, request.Name)

	if len(duplicatedUser.ID) != 0 {
		// User already exists
		return nil, &exception.DuplicateEntryError{
			Message: fmt.Sprintf("%v is already in the database !", request.Name),
		}
	}

	// Create new user
	userEntity := &domain.User{
		ID:        "696969",
		Name:      request.Name,
		Age:       request.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert to db !
	created, err := this.UserRepository.Create(this.Database, userEntity)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return created, nil
}
