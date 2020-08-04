package main

import (
	"log"
	"strings"
)

// checks and returns correct args

const ARGSDELIMITER = "-"

func parseArgs(args []string, config *ProgramConfig) {

	log.Println(args)
	// reads all arguments and finds argument delimiter
	for delimiterIndex, val := range args {
		if strings.Contains(val, ARGSDELIMITER) {
			config.listFormerCommand = TerminalCommand{args[0], args[1:delimiterIndex]}

			secondCommandBeginning := delimiterIndex+1
			config.executableCommand = TerminalCommand{args[secondCommandBeginning], args[secondCommandBeginning+1:]}

			log.Println(config.listFormerCommand.commandName, config.listFormerCommand.commandArgs)
			log.Println(args[secondCommandBeginning], args[secondCommandBeginning+1:])
			break
		}
	}
}
