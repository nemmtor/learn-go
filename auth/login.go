package auth

import (
	"fmt"
)

const validLogin = "Kacper"
const validPassword = "password123"
const maxTries = 3

// Login authorizes user
func Login() {
	var username string
	var password string

	fmt.Println("Witaj w systemie logowania.")

	try := 0

	for ; try < maxTries; try++ {
		fmt.Println("Podaj login:")
		fmt.Scanln(&username)

		fmt.Println("Podaj hasło:")
		fmt.Scanln(&password)

		if (username != validLogin) || (password != validPassword) {
			fmt.Println("Wpisałeś zły login lub hasło.")
		} else {
			fmt.Println("Logowanie się powiodło.")
			break
		}
	}

	if try == maxTries {
		fmt.Printf("Konto zablokowane. Wpisałeś %d razy złe hasło", try)
	}
}
