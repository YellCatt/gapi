package router

import (
	"net/http"

	"github.com/example/gapi/controller"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/swaggo/swag"
)

func NewRouter(userController *controller.UserController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", userController.CreateUser)
	mux.HandleFunc("GET /api/users", userController.GetAllUsers)
	mux.HandleFunc("GET /api/users/{id}", userController.GetUserByID)
	mux.HandleFunc("PUT /api/users/{id}", userController.UpdateUser)
	mux.HandleFunc("DELETE /api/users/{id}", userController.DeleteUser)

	mux.HandleFunc("GET /swagger/", func(w http.ResponseWriter, r *http.Request) {
		httpSwagger.WrapHandler(w, r, httpSwagger.URL("/swagger/doc.json"))
	})

	mux.HandleFunc("GET /swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(swag.GetSwagger()))
	})

	return mux
}