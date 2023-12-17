package controllers

import (
	"encoding/json"
	"log"
)

func GetJson(object interface{}) []byte {
	json, err := json.Marshal(object)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return json
}
