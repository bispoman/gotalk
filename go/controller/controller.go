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

	result := service.GetPeoplefromSwapi(id)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, result)
}
