package database

import (
	"encoding/json"
	"errors"
	"learn-go/jsonio"
	"learn-go/utils"
	"log"
)

var store = &Store

func init() {
	log.Println("Opening db file...")
}

// Connect reads json database
func Connect() error {
	if !utils.PathExists(dbFilename) {
		jsonio.CreateNewFile(dbFilename)
	}

	jsonData, error := jsonio.GetJSONData(dbFilename)
	if error != nil {
		logCorruptedFile()
		utils.MoveFile(dbFilename, corruptedFilesPath)
	}

	jsonError := json.Unmarshal(jsonData, store)
	if jsonError != nil {
		logCorruptedFile()
		utils.MoveFile(dbFilename, corruptedFilesPath)
		return errors.New("Couldn't connect to database file")
	}
	log.Println("Connected to database file.")
	return nil
}

func logCorruptedFile() {
	log.Printf("File is corrupted, moving it to %v directory and creating a new one...\n", corruptedFilesPath)
}
