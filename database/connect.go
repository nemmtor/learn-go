package database

import (
	"encoding/json"
	"io/ioutil"
	"learn-go/utils"
	"log"
	"os"
)

var databaseObject = &Store

func init() {
	log.Println("Opening db file...")
}

// Connect reads json database
func Connect(filename string) {
	jsonData := getJSONData(filename)

	jsonError := json.Unmarshal(jsonData, databaseObject)
	if jsonError != nil {
		log.Println("File is corrupted, moving it to corrupted-files and creating a new one...")
		moveCorruptedFile(filename)
		Connect(filename)
	}
	log.Println("Connected to database file.")
}

func saveJSON(jsonData []byte) {
	error := ioutil.WriteFile("db.json", jsonData, 0644)
	if error != nil {
		log.Fatal(error)
	}
}

func getJSONData(filename string) []byte {
	jsonFile, fileOpenError := os.Open(filename)
	if fileOpenError != nil {
		log.Println("Couldn't open db file, creating one...")
		jsonFile = createNewFile(filename)
	}

	jsonBytes, bytesReadError := ioutil.ReadAll(jsonFile)
	if bytesReadError != nil {
		moveCorruptedFile(filename)
	}

	defer jsonFile.Close()
	return jsonBytes
}

func createNewFile(filename string) *os.File {
	jsonFile, error := os.Create(filename)

	var newJSON struct{}
	jsonProduct, jsonProductError := json.Marshal(newJSON)
	writeFileError := ioutil.WriteFile(filename, jsonProduct, 0644)

	if error != nil || jsonProductError != nil || writeFileError != nil {
		log.Fatal("Couldn't create db file, exiting.")
	}

	return jsonFile
}

func moveCorruptedFile(filename string) {
	dirName := "corrupted-files/"

	if !utils.PathExists(dirName) {
		if os.Mkdir(dirName, 0775) != nil {
			log.Println(os.Mkdir(dirName, 0775))
			log.Fatal("Couldn't create directory, exiting.")
		}
	}
	newFilename := dirName + utils.GetCurrentDatetime() + "json"
	if os.Rename(filename, newFilename) != nil {
		log.Fatal("Couldn't move file, exiting.")
	}
}
