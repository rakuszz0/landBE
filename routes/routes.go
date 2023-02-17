package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	TrainRoutes(r)
	StasiunRoutes(r)
	TiketRoutes(r)
	TransactionRoutes(r)
}
