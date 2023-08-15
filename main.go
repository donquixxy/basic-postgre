package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"postgre-basic/config"
	"postgre-basic/database/postgre"
	"postgre-basic/database/redis"
	"postgre-basic/internal/api/routes"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/handler"
	"postgre-basic/internal/repository"
	"postgre-basic/internal/usecases"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	appConfig := config.GetConfig()

	log.Println("This App is running in the config :", appConfig.Application.Server)

	db, err := postgre.InitPostgreDb(*appConfig.Database)
	resisClient := redis.NewRedisClient()
	if err != nil {
		log.Fatalln("Error opening postgre database !", err.Error())
	}
	e := echo.New()
	e.HTTPErrorHandler = exception.ErrorRouteHandler

	go func() {
		err := e.Start(":8181")

		if err != nil {
			log.Fatalln("Error starting echo", err.Error())
		}
	}()

	// Repository Layer
	userRepository := repository.NewUserRepository()
	companyRepository := repository.NewCompanyRepository()
	// End Of Repository Layer

	// Usecases layer
	userServices := usecases.NewUserServices(db, userRepository, resisClient)
	companyServices := usecases.NewCompanyServices(companyRepository, db)
	// End of Usecase Layer

	// Handler layer
	userHandler := handler.NewUserHandler(userServices, resisClient)
	companyHandler := handler.NewCompanyHandler(companyServices)
	// End of Handler Layer

	routes.UserRoutes(e, userHandler)
	routes.CompanyRoute(e, companyHandler)

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

	postgre.ClosePostgreConnection(db)

}
