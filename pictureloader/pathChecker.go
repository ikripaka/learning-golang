package main

import (
	"errors"
	"fmt"
	"os"
)

// checks whether the path is exists

func IsPathsCorrect(filePath, folderPath string) (bool, error) {
	fmt.Println(filePath, folderPath)

	// checks if file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, errors.New("file doesn't exist")
	}

	// checks if folder exists
	_, err = os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false, errors.New("folder doesn't exist")
	}

	return true, nil
}
