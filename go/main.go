package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bispoman/gotalk/router"
)

func main() {

	router := router.GetRoutes()

	http.Handle("/", router)

	port := ":8080"

	log.Info("Application has been started at port ", port)

	log.Fatal(http.ListenAndServe(port, nil))

}
