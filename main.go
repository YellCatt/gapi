package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example/gapi/config"
	"github.com/example/gapi/controller"
	"github.com/example/gapi/repository"
	"github.com/example/gapi/router"
	"github.com/example/gapi/service"
)

func main() {
	config.LoadConfig()

	db := config.NewDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := router.NewRouter(userController)

	port := config.GetServerPort()
	log.Printf("Server starting on :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}