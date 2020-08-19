package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// checks and returns correct args

// max arguments number is 2 (file path value/folder path value)
const MAXARGSNUMBER = 1

// Reads file with urls and pushes them to the buffered channel
// filePath - filepath for file with urls
// config - program configuration that contains all variables that need

func ReadConfig(filePath string) (config *ProgramConfig) {
	// reads all file in byte -> convert it to string -> splits it with new line symbol -> splits line with delimiter
	config = &ProgramConfig{}
	allFileInByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Problems with reading file")
	}
	sliceOfCommandAssociations := strings.Split(string(allFileInByte), "\n")
	config.listOfSelectableItems = make([]TerminalCommand, len(sliceOfCommandAssociations))
	for i, val := range sliceOfCommandAssociations {
		splittedData := strings.Split(val, "::")
		config.listOfSelectableItems[i] = TerminalCommand{splittedData[0], splittedData[len(splittedData)-1]}
	}
	return
}

// Checks whether the path is exists
func IsPathsCorrect(filePath string) error {

	// checks if file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return INCORRECTFILEPATH
	}

	return nil
}

// Checks and returns correct args
func getArgs(args []string) (filePath string) {
	//checks if number of args in slice is correct
	if len(args) != MAXARGSNUMBER {
		log.Fatal("Incorrect number of arguments ", "requires:", MAXARGSNUMBER, " has:", len(args))
	}

	filePath = args[0]
	// checks if folder/file path exists
	err := IsPathsCorrect(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return
}
