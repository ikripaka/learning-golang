package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
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
	if strings.Contains(config.listOfSelectableItems[config.userChoiceListIndex].command, " && ") {
		splittedCommands := strings.Split(config.listOfSelectableItems[config.userChoiceListIndex].command, " && ")
		for _, val := range splittedCommands {
			execCommand(val)
		}
	} else {
		execCommand(config.listOfSelectableItems[config.userChoiceListIndex].command)
	}

}
func execCommand(command string) {
	parsedCommand := strings.Split(command, " ")

	cmd := exec.Command(parsedCommand[0], parsedCommand[1:]...)
	//redirected the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// channel to handle ctrl+c
	c := make(chan os.Signal)
	// handling SIGTERM
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// func that processing SIGTERM
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
	}()

}
