package routes

import (
	"landtick/handlers"
	"landtick/pkg/connection"
	"landtick/pkg/middleware"
	"landtick/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transRepository := repositories.RepositoryTrans(connection.DB)
	h := handlers.HandlerTransaction(transRepository)

	r.HandleFunc("/transall", middleware.Auth(h.FindsTransaction)).Methods("GET")
	r.HandleFunc("/transactions", middleware.Auth(h.TransactionUser)).Methods("GET")
	r.HandleFunc("/detailtransaction/{id}", middleware.Auth(h.GetTrans)).Methods("GET")
	r.HandleFunc("/transpending", middleware.Auth(h.TransactionPending)).Methods("GET")
	r.HandleFunc("/transactionpayment/{id}", middleware.Auth(h.Payment)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")

}
