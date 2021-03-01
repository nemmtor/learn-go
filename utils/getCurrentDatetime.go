package utils

import "time"

// GetCurrentDatetime returns current datetime in a readable string
func GetCurrentDatetime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}
