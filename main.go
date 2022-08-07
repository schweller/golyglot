package main

import (
	"os"

	cli "github.com/schweller/golyglot/command"
)

func main() {
	cli.CreateRootCommand().Execute()
	os.Exit(1)
}
