package user

import (
	"learn-go/database"
)

// Repository for User Entity
type Repository struct {
}

// CreateUser saves user into database
func (repository Repository) CreateUser(user database.User) {
	database.Store.Users = append(database.Store.Users, user)
	database.Store.Save()
}
