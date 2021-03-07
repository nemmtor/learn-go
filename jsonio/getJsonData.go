package jsonio

import (
	"io/ioutil"
	"log"
	"os"
)

// GetJSONData ...
func GetJSONData(filename string) ([]byte, error) {
	jsonFile, error := os.Open(filename)

	if error != nil {
		log.Fatal("Couldn't open db file, exiting...")
	}

	jsonBytes, bytesReadError := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()
	return jsonBytes, bytesReadError
}
