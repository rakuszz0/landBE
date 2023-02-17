package routes

import (
	"landtick/handlers"
	"landtick/pkg/connection"
	"landtick/pkg/middleware"
	"landtick/repositories"

	"github.com/gorilla/mux"
)

func TrainRoutes(r *mux.Router) {
	trainRepository := repositories.RepositoryTrain(connection.DB)
	h := handlers.HandlerTrain(trainRepository)

	r.HandleFunc("/trains", h.FindTrains).Methods("GET")
	r.HandleFunc("/train/{id}", h.GetTrain).Methods("GET")
	r.HandleFunc("/train/{id}", middleware.Auth(h.DeleteTrain)).Methods("DELETE")
	r.HandleFunc("/train/{id}", middleware.Auth(h.UpdateTrain)).Methods("PATCH")
	r.HandleFunc("/train", middleware.Auth(h.CreateTrain)).Methods("POST")
}
