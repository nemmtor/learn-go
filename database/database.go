package database

import "learn-go/auth"

// Database implemented in JSON
type Database struct {
	Users []auth.User `json:"users"`
}
