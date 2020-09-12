package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/liampulles/countdown/pkg/command"
)

var registeredCommands = map[string]command.Command{
	"conundrum": command.Conundrum,
	"letters":   command.Letters,
}

// Run is a delegate for the main function,
// to abstract os calls and expose a public method for testability.
func Run(args []string) int {
	if len(args) < 2 {
		logError("You must specify a command. Valid commands: [%s]\n", validCommands())
		return 1
	}

	commandName := args[1]
	command, ok := registeredCommands[commandName]
	if !ok {
		logError("%s is not a valid command. Valid commands: [%s]\n", commandName, validCommands())
		return 2
	}

	commandArgs := args[2:]
	if err := command(commandArgs); err != nil {
		logError("ERROR: [%v]\n", err)
		return 3
	}

	return 0
}

func validCommands() string {
	var list []string
	for k := range registeredCommands {
		list = append(list, k)
	}
	return strings.Join(list, ", ")
}

func logError(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
