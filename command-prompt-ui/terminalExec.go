package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
	"strings"
)

func formList(config *ProgramConfig){
	log.Println(config.listFormerCommand.commandName, config.listFormerCommand.commandArgs)
	// execute list former command
	cmd:=exec.Command(config.listFormerCommand.commandName, config.listFormerCommand.commandArgs...)
	b, err := cmd.CombinedOutput()
	if err != nil{
		log.Println("54")
		log.Fatal(err)
	}
	log.Println(string(b))

	// split command output by "\n"
	config.listFormerOutput = strings.Split(string(b), "\n")

	prompt := promptui.Select{
		Label: "Select settings for the command execution",
		Items: config.listFormerOutput,
	}
	_, result, err := prompt.Run()
	fmt.Println("Result: ", result)
}