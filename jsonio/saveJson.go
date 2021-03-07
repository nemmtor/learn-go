package jsonio

import (
	"io/ioutil"
	"log"
)

// SaveJSON saves json input into given filename
func SaveJSON(jsonData []byte, filename string) {
	error := ioutil.WriteFile(filename, jsonData, 0644)
	if error != nil {
		log.Fatal(error)
	}
}
