package usecases

import (
	"fmt"
	"postgre-basic/internal/api/request/userrequest"
	"postgre-basic/internal/domain"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/repository"
	"postgre-basic/utils/etc"
	"time"

	"gorm.io/gorm"
)

type UserServices interface {
	CreateUser(request *userrequest.CreateRequest) (*domain.User, error)
	FindAllUsers() ([]*domain.User, error)
	UpdateUsers(request *userrequest.UpdateRequest, id string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	Delete(id string) error
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
		ID:        etc.GenerateRandomUUID(),
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

func (this *UserServicesImpl) FindAllUsers() ([]*domain.User, error) {
	listUsers := this.UserRepository.FindAll(this.Database)

	if len(listUsers) == 0 {
		return nil, &exception.RecordNotFoundError{
			Message: "No users found",
		}
	}

	return listUsers, nil
}

func (this *UserServicesImpl) UpdateUsers(request *userrequest.UpdateRequest, id string) (*domain.User, error) {
	// Find user by id first
	currentUser, err := this.UserRepository.FindByID(this.Database, id)

	if err != nil {
		return nil, &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}

	// Update the value
	currentUser.Age = request.Age
	currentUser.Name = request.Name

	// Update to the db !
	err = this.UserRepository.Update(this.Database, currentUser)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	// Get the newest data for API return data
	newData, _ := this.UserRepository.FindByID(this.Database, id)

	return newData, nil
}

func (this *UserServicesImpl) FindByID(id string) (*domain.User, error) {
	// Find user by id first
	currentUser, err := this.UserRepository.FindByID(this.Database, id)

	if err != nil {
		return nil, &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}
	return currentUser, nil
}

func (this *UserServicesImpl) Delete(id string) error {
	// Find user by id first
	currentUser, err := this.UserRepository.FindByID(this.Database, id)

	if err != nil {
		return &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}

	errDelete := this.UserRepository.Delete(this.Database, currentUser)

	if errDelete != nil {
		return &exception.BadRequestError{Message: errDelete.Error()}
	}

	return nil
}
