package user

import "fmt"

// Repository for User Entity
type Repository struct {
}

// CreateUser saves user into database
func (repository Repository) CreateUser(user Entity) {
	fmt.Println(user)
}
