package database

import "learn-go/user"

// Store content
var Store = database{}

type database struct {
	Users []user.Entity `json:"users"`
}
