package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"postgre-basic/config"
	"postgre-basic/database/postgre"
	"postgre-basic/internal/api/routes"
	"postgre-basic/internal/handler"
	"postgre-basic/internal/repository"
	"postgre-basic/internal/usecases"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello world")

	appConfig := config.GetConfig()

	fmt.Println("Name :", appConfig.Application.Name)
	fmt.Println("pass :", appConfig.Database.Password)

	db, err := postgre.InitPostgreDb()

	if err != nil {
		log.Fatalln("Error opening postgre database !", err.Error())
	}

	e := echo.New()

	go func() {
		err := e.Start(":8181")

		if err != nil {
			log.Fatalln("Error starting echo", err.Error())
		}
	}()

	// Repository Layer
	userRepository := repository.NewUserRepository()
	// End Of Repository Layer

	// Usecases layer
	userServices := usecases.NewUserServices(db, userRepository)
	// End of Usecase Layer

	// Handler layer
	userHandler := handler.NewUserHandler(userServices)
	// End of Handler Layer

	routes.UserRoutes(e, userHandler)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
