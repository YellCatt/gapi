package main

import (
	"log"
	"net/http"

	"github.com/example/gapi/config"
	"github.com/example/gapi/controller"
	"github.com/example/gapi/repository"
	"github.com/example/gapi/router"
	"github.com/example/gapi/service"
)

func main() {
	db := config.NewDatabase()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := router.NewRouter(userController)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}