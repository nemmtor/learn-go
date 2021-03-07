package database

type database struct {
	Users []User `json:"users"`
}

// User entity
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
