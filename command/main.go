package command

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

var (
	authKey    string
	sourceLang string
	targetLang string
)

func Run(args []string) int {
	return RunCustom(args)
}

func RunCustom(args []string) int {
	commands := initCommands()

	cli := &cli.CLI{
		Name:     "deepl",
		Args:     args,
		Commands: commands,
	}
	exitCode, err := cli.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
