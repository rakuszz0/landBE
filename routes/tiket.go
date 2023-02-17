package routes

import (
	"landtick/handlers"
	"landtick/pkg/connection"
	"landtick/pkg/middleware"
	"landtick/repositories"

	"github.com/gorilla/mux"
)

func TiketRoutes(r *mux.Router) {
	tiketRepository := repositories.RepositoryTiket(connection.DB)
	h := handlers.HandlerTiket(tiketRepository)

	r.HandleFunc("/tiket", middleware.Auth(h.CreateTiket)).Methods("POST")
	r.HandleFunc("/search-tiket", h.FilterTiket).Methods("POST")
	r.HandleFunc("/tikets", h.GetAllTiket).Methods("GET")
	r.HandleFunc("/tiket-transaction/{id}", middleware.Auth(h.CreateTransTiket)).Methods("POST")
	r.HandleFunc("/tiket/{id}", middleware.Auth(h.DeleteTiket)).Methods("DELETE")

}
