package postgre

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreConfig struct {
	Driver       string
	Username     string
	Password     string
	Port         uint
	Address      string
	DatabaseName string
	Schemas      string
	SSLMode      string
}

func InitPostgreDb(config PostgreConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=agusari password=12345678 sslmode=disable", config.Address)
	database, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		fmt.Println("error opening database", er.Error())
		return nil, er
	}

	db, er := database.DB()

	if er != nil {
		return nil, er
	}

	er = db.Ping()

	if er != nil {
		return nil, er
	}

	fmt.Println("Connected to database!")

	return database, nil
}
