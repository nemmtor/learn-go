package user

import (
	"errors"
	"learn-go/database"
	"log"

	"github.com/google/uuid"
)

// Repository for User Entity
type Repository struct{}

// Create saves user into database
func (repository Repository) Create(user database.User) {
	_, error := repository.GetByUsername(user.Username)
	if error == nil {
		log.Printf("Username %v already exists.", user.Username)
		return
	}

	newUser := user
	newUser.ID = uuid.NewString()

	database.Store.Users = append(database.Store.Users, newUser)
	database.Store.Save()
}

// GetByUsername returns user entity or error
func (repository Repository) GetByUsername(username string) (database.User, error) {
	for i := range database.Store.Users {
		if database.Store.Users[i].Username == username {
			return database.Store.Users[i], nil
		}
	}
	return database.User{}, errors.New("User not found")
}
