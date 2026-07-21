package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example/gapi/config"
	"github.com/example/gapi/controller"
	"github.com/example/gapi/logger"
	"github.com/example/gapi/repository"
	"github.com/example/gapi/router"
	"github.com/example/gapi/service"

	"go.uber.org/zap"
)

func main() {
	config.LoadConfig()

	if err := config.InitDirectories(); err != nil {
		log.Fatalf("failed to init directories: %v", err)
	}

	if err := logger.Init(config.GetLogPath(), config.GetLogLevel()); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	db := config.NewDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := router.NewRouter(userController)

	port := config.GetServerPort()
	logger.Info("server starting", zap.Int("port", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		logger.Fatal("server exited", zap.Error(err))
	}
}