package main

import (
	"os"

	"github.com/schweller/rumor/command"
)

func main() {
	command.CreateRootCommand().Execute()
	os.Exit(1)
}
