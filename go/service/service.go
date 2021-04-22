package service

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetPeoplefromSwapi(id string) string {
	log.Info("requesting to swapi people ", id)
	url := "https://swapi.dev/api/people/" + id
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	log.Info("Request to swapi/people made successfully")
	return responseString
}
