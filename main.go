package main

import (
	"learn-go/database"
	"learn-go/user"
	"log"
)

func main() {
	database.Connect("db.json")
	// Fake register
	newUser := database.User{Username: "Kacper", Password: "123"}

	userRepository := user.Repository{}
	userRepository.CreateUser(newUser)
	log.Println(database.Store)
}
