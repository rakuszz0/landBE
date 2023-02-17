package routes

import (
	"landtick/handlers"
	"landtick/pkg/connection"
	"landtick/pkg/middleware"
	"landtick/repositories"

	"github.com/gorilla/mux"
)

func StasiunRoutes(r *mux.Router) {
	stasiunRepository := repositories.RepositoryStasiun(connection.DB)
	h := handlers.HandlerStasiun(stasiunRepository)

	r.HandleFunc("/stasiuns", h.FindStasiuns).Methods("GET")
	r.HandleFunc("/stasiun/{id}", h.GetStasiun).Methods("GET")
	r.HandleFunc("/stasiun/{id}", middleware.Auth(h.DeleteStasiun)).Methods("DELETE")
	r.HandleFunc("/stasiun/{id}", middleware.Auth(h.UpdateStasiun)).Methods("PATCH")
	r.HandleFunc("/stasiun", middleware.Auth(h.CreateStasiun)).Methods("POST")
}
