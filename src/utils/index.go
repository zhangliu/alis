package utils

import (
	"os"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IsFileExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, err }
	return true, err
}