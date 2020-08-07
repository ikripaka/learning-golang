package main

import (
	"errors"
	"os"
)

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
	userChoice            string
	userChoiceListIndex   int
}

//this struct represents terminal command separation
type TerminalCommand struct {
	listNaming string
	command    string
}

func (command *TerminalCommand) String() string {
	return command.listNaming + ": " + command.command
}

func main() {
	config := ProgramConfig{}

	filepath := getArgs(os.Args[1:])
	ReadListConfig(filepath, &config)
	formList(&config)
	executeSelectedCommand(&config)
}
