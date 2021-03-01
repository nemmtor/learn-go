package main

import (
	"learn-go/database"
	"learn-go/user"
	"log"
)

func main() {
	database.Connect("db.json")
	log.Println(database.Store)

	// Fake register
	newUser := user.Entity{Username: "Kacper", Password: "123"}

	userRepository := user.Repository{}
	userRepository.CreateUser(newUser)
}
