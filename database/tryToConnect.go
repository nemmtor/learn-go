package database

var connectionMaxTries = 2
var triesCounter int

func init() {
	triesCounter = 1
}

// TryToConnect tries to connect to db file
func TryToConnect() {
	connectionError := Connect()
	for connectionError != nil && triesCounter <= connectionMaxTries {
		connectionError = Connect()
	}
}
