package main

import (
	"errors"
	"os"
)

// This program is realisation of Command Prompt UserInterface

// Firstly to run it you should write your config file in this way " command name/description::command"
// Then build it and run with one argument - absolute path to the file

// error that occurs with filepath
var INCORRECTFILEPATH = errors.New("file doesn't exist")

// error that occurs with library execution forming list
type LISTEXECUTIONERROR struct {
	err error
}

func (e *LISTEXECUTIONERROR) Error() string {
	return "problems in mainfoldco/promptui execution, error: " + e.err.Error()
}

// this struct helps to contain all program configuration variable at one place
type ProgramConfig struct {
	listOfSelectableItems []TerminalCommand
}

//this struct represents terminal command separation
type TerminalCommand struct {
	listNaming string
	command    string
}

func main() {
	filepath := getArgs(os.Args[1:])

	// reading configuration file
	config := ReadConfig(filepath)

	// forming list which return user choice
	listIndex, _ := formList(config.listOfSelectableItems)

	// executing chosen command
	executeSelectedCommand(config.listOfSelectableItems[listIndex].command)
}
