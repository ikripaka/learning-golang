package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// max arguments number is 2 (file path value/folder path value)
const MAXARGSNUMBER = 2

// Reads file with urls and pushes them to the buffered channel
// filePath - filepath for file with urls
// config - program configuration that contains all variables that need

func ReadPictureUrls(filePath string, config *ProgramConfig) {
	// reads all file in byte -> convert it to string -> splits it with new line symbol

	allFileInByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Problems with reading file")
	}
	sliceOfUrls := strings.Split(string(allFileInByte), "\n")
	config.pictureUrlsChan = make(chan Item, len(sliceOfUrls))
	for _, val := range sliceOfUrls {
		config.pictureUrlsChan <- Item{url: val}
	}
	close(config.pictureUrlsChan)
}

// checks whether the path is exists

func IsPathsCorrect(filePath, folderPath string) error {

	// checks if file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return errors.New("file doesn't exist")
	}

	// checks if folder exists
	_, err = os.Stat(folderPath)
	if os.IsNotExist(err) {
		return errors.New("folder doesn't exist")
	}

	return nil
}

// checks and returns correct args

func getArgs(args []string) (urlFilePath, folderPath string) {
	//checks if number of args in slice is correct
	if len(args) != MAXARGSNUMBER {
		log.Fatal("Incorrect number of argumets ", "requires:", MAXARGSNUMBER, " has:", len(args))
	}

	urlFilePath = args[0]
	folderPath = args[1]
	// checks if folder/file path exists
	err := IsPathsCorrect(urlFilePath, folderPath)
	if err != nil {
		log.Fatal(err)
	}

	return
}
