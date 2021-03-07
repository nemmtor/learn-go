package jsonio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var errorMsg = "Couldn't create db file, exiting"

// CreateNewFile ...
func CreateNewFile(filename string) *os.File {
	file, error := os.Create(filename)

	if error != nil {
		log.Fatal(errorMsg)
	}

	writeEmptyJSON(filename)

	return file
}

func writeEmptyJSON(filename string) {
	var newJSON struct{}

	bytes, error := json.Marshal(newJSON)
	if error != nil {
		log.Fatal(errorMsg)
	}

	error = ioutil.WriteFile(filename, bytes, 0644)

	if error != nil {
		log.Fatal(errorMsg)
	}
}
