package database

type database struct {
	Users []User `json:"users"`
}

// User entity
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
