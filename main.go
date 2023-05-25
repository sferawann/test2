package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sferawann/test2/config"
	"github.com/sferawann/test2/controller"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/repository"
	"github.com/sferawann/test2/router"
	"github.com/sferawann/test2/service"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	// db.Table("borrowers").AutoMigrate(&model.Users{})
	db.Table("borrowers").Find(&model.Borrower{})
	db.Table("lenders").Find(&model.Lender{})

	//Init Repository
	borRepo := repository.NewBorrowerRepositoryImpl(db)
	lenRepo := repository.NewLenderRepositoryImpl(db)

	//Init Service
	authService := service.NewAuthServiceImpl(borRepo, validate)
	borService := service.NewBorrowerServiceImpl(borRepo, validate)
	lenService := service.NewLenderServiceImpl(lenRepo, validate)

	//Init controller
	authController := controller.NewAuthController(authService)
	borController := controller.NewBorrowerController(borService)
	lenController := controller.NewLenderController(lenService)

	routes := router.NewRouter(borRepo, authController, borController,lenController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
