package main

import "os"

// this struct helps to contain all program configuration variable at one place
type ProgramConfig struct {
	listFormerCommand TerminalCommand
	executableCommand TerminalCommand

	listFormerOutput []string
	userChoice string
}

//this struct represents terminal command separation
type TerminalCommand struct {
	commandName string
	commandArgs []string
}

func main() {
	config := ProgramConfig{}
	parseArgs(os.Args[1:], &config)

	formList(&config)


}
