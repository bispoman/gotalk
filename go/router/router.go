package router

import (
	"github.com/bispoman/gotalk/controller"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {

	routes := mux.NewRouter().StrictSlash(false)

	routes = routes.PathPrefix("/api/v1").Subrouter()

	routes.HandleFunc("/healthcheck", controller.Healthcheck).Methods("GET")

	routes.HandleFunc("/getpeople/:id", controller.GetPeople).Methods("GET")

	return routes
}
