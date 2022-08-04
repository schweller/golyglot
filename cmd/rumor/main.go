package main

import (
	"os"

	"github.com/schweller/rumor/command"
)

func main() {
	command.CreateCommand().Execute()
	os.Exit(1)
}
