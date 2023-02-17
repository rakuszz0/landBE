package routes

import (
	"landtick/handlers"
	"landtick/pkg/connection"
	"landtick/pkg/middleware"
	"landtick/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(connection.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/checkauth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
