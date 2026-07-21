package router

import (
	"net/http"

	"github.com/example/gapi/controller"
)

func NewRouter(userController *controller.UserController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", userController.CreateUser)
	mux.HandleFunc("GET /api/users", userController.GetAllUsers)
	mux.HandleFunc("GET /api/users/{id}", userController.GetUserByID)
	mux.HandleFunc("PUT /api/users/{id}", userController.UpdateUser)
	mux.HandleFunc("DELETE /api/users/{id}", userController.DeleteUser)

	return mux
}