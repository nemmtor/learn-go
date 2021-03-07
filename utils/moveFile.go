package utils

import (
	"log"
	"os"
)

// MoveFile moves given file to new directory
func MoveFile(filename string, path string) {
	if !PathExists(path) {
		if os.Mkdir(path, 0775) != nil {
			log.Println(os.Mkdir(path, 0775))
			log.Fatal("Couldn't create directory, exiting.")
		}
	}
	newFilename := path + GetCurrentDatetime() + ".json"
	if os.Rename(filename, newFilename) != nil {
		log.Fatal("Couldn't move file, exiting.")
	}
}
