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

func InitPostgreDb() (*gorm.DB, error) {
	dsn := "host=localhost user=agusari password=12345678 dbname=postgres port=32768 sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
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
