package main

import (
	"os"

	cli "github.com/schweller/golyglot/command"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
