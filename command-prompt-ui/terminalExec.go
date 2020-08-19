package main

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
)

// Command beginning for command execution with pipe (|)
var DEFAULTEXECUTIONCOMMAND = []string{"sh", "-c"}

// Forms list depending on user file input
func formList(listOfSelectableItems []TerminalCommand) (listIndex int, userChoice string) {
	// creates lists
	listNames := make([]string, len(listOfSelectableItems))
	listOfCommands := make([]string, len(listOfSelectableItems))

	//parsing commands name and args
	for i, val := range listOfSelectableItems {
		listNames[i] = val.listNaming
		listOfCommands[i] = val.command
	}

	//creating list in cli
	prompt := promptui.Select{
		Label: "Select settings for the command execution",
		Items: listNames,
		Size:  len(listNames),
	}

	//running selection in command prompt
	listIndex, userChoice, err := prompt.Run()
	if err != nil {
		log.Fatal(LISTEXECUTIONERROR{err: err})
	}

	return
}

// Executing selected command using command beginning DEFAULTEXECUTIONCOMMAND
func executeSelectedCommand(command string) {
	cmd := exec.Command(DEFAULTEXECUTIONCOMMAND[0], DEFAULTEXECUTIONCOMMAND[1], command)
	//redirected the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// running command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
