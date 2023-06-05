package main

import (
	"fmt"
	"log"
	"postgre-basic/config"
	"postgre-basic/database/postgre"
)

func main() {
	fmt.Println("Hello world")

	appConfig := config.GetConfig()

	fmt.Println("Name :", appConfig.Application.Name)
	fmt.Println("pass :", appConfig.Database.Password)

	_, err := postgre.InitPostgreDb(*appConfig.Database)

	if err != nil {
		log.Fatalln("Error opening postgre database !", err.Error())
	}

}
