package usecases

import (
	"context"
	"fmt"
	"postgre-basic/internal/api/request/userrequest"
	"postgre-basic/internal/domain"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/repository"
	"postgre-basic/utils/etc"
	"time"

	"github.com/redis/go-redis/v9"
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
	RedisClient    *redis.Client
}

func NewUserServices(
	Database *gorm.DB,
	UserRepository repository.UserRepository,
	RedisClient *redis.Client,
) UserServices {
	return &UserServicesImpl{
		Database:       Database,
		UserRepository: UserRepository,
		RedisClient:    RedisClient,
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
		IDCompany: "15e59302-b227-4c04-8889-a093ebfe1a68",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Test create to redis !
	err := this.RedisClient.Set(context.TODO(), userEntity.ID, userEntity, time.Minute*5).Err()

	if err != nil {
		fmt.Println("Redis client error create data :", err.Error())
		return nil, err
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

	// Delete redis value
	erro := this.RedisClient.Del(context.TODO(), id).Err()

	if erro != nil {
		fmt.Println("Error deleting redis value")
		fmt.Println(erro.Error())
	}

	// Create new value key redis
	erro = this.RedisClient.Set(context.TODO(), id, currentUser, time.Minute*5).Err()

	if erro != nil {
		fmt.Println("Error creating new redis value")
		fmt.Println(erro.Error())
	}

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

	// Find data at cahce first
	cachedData, err := this.RedisClient.Get(context.TODO(), id).Result()

	// If not found
	if err != nil {
		// Search data at db
		// Find user by db
		currentUser, err := this.UserRepository.FindByID(this.Database, id)
		if err != nil {
			return nil, &exception.RecordNotFoundError{
				Message: err.Error(),
			}
		}

		// Set data to cache !
		this.RedisClient.Set(context.TODO(), currentUser.ID, currentUser, time.Minute*5)
		fmt.Println("Data were found at db")
		GetAllRedisKeys(this.RedisClient)
		return currentUser, nil
	}

	// If found at cache
	// Unmarshal to object
	payload := &domain.User{}

	payload.UnmarshalBinary([]byte(cachedData))
	fmt.Println("Data were found at cache")
	GetAllRedisKeys(this.RedisClient)
	return payload, nil

}

func GetAllRedisKeys(c *redis.Client) {
	listKeys, err := c.Keys(context.TODO(), "*").Result()

	if err != nil {
		fmt.Println(err)
	}

	for _, key := range listKeys {
		fmt.Println(key)
	}
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
