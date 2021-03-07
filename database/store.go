package database

import (
	"encoding/json"
	"log"
)

// Store content
var Store = database{}

func (store database) Save() {
	jsonData, jsonDataError := json.Marshal(Store)
	if jsonDataError != nil {
		log.Fatal("Could'nt convert store to json.")
	}
	saveJSON(jsonData)
}
