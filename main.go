package main

import (
	"learn-go/database"
	"learn-go/user"
	"log"
)

func main() {
	database.TryToConnect()

	newUser := database.User{Username: "Kacper", Password: "123"}
	userRepository := user.Repository{}
	userRepository.Create(newUser)

	foundUser, error := userRepository.GetByUsername("Michał")
	if error != nil {
		log.Println(error)
	} else {
		log.Printf("Found user with username: %v\n", foundUser.Username)
	}
}
