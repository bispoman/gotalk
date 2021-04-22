package service

import (
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func GetPeoplefromSwapi(id string) string {
	log.Info("requesting people %s to swapi", id)
	response, err := http.Get("https://swapi.dev/api/people/" + id)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	responseString := string(responseData)
	log.Info("Request to swapi/people made successfully, result: ", responseData)
	return responseString
}
