package utils

import "os"

// PathExists checks if given path exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
