package main

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"strings"
)

// forms list depending on user file input
func formList(config *ProgramConfig) {
	// creates lists
	listNames := make([]string, len(config.listOfSelectableItems))
	listOfCommands := make([]string, len(config.listOfSelectableItems))

	//parsing commands name and args
	for i, val := range config.listOfSelectableItems {
		listNames[i] = val.listNaming
		listOfCommands[i] = val.command
	}

	//creating list in cli
	prompt := promptui.Select{
		Label: "Select settings for the command execution",
		Items: listNames,
		Size:  cap(listNames),
	}

	//running selection in command prompt
	index, result, err := prompt.Run()
	if err != nil {
		log.Fatal(LISTEXECUTIONERROR{err: err})
	}
	config.userChoice = result
	config.userChoiceListIndex = index

}

// executing selected command
func executeSelectedCommand(config *ProgramConfig) {
	//splitting command by " "
	parsedCommand := strings.Split(config.listOfSelectableItems[config.userChoiceListIndex].command, " ")

	cmd := exec.Command(parsedCommand[0], parsedCommand[1:]...)
	//redirected the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
