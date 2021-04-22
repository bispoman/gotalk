package controller

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/bispoman/gotalk/service"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Info("healthcheck hit, responding with ok")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Ok")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Info("getPeople hit, forwarding to service")

	result := service.GetPeoplefromSwapi(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, result)
}
