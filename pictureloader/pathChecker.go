package main

import (
	"errors"
	"fmt"
	"os"
)

// checks whether the path is exists

func IsPathsCorrect(filePath, folderPath string) (bool, error) {
	fmt.Println(filePath, folderPath)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, errors.New("file doesn't exist")
	}

	_, err = os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false, errors.New("folder doesn't exist")
	}
	return true, nil
}
