package command

import (
	"os"

	"github.com/mitchellh/cli"
)

func initCommands() map[string]cli.CommandFactory {
	commands := map[string]cli.CommandFactory{
		"translate": func() (cli.Command, error) {
			return &TranslateCommand{
				UI: &cli.BasicUi{
					Reader: os.Stdin,
					Writer: os.Stdout,
				},
			}, nil
		},
		"usage": func() (cli.Command, error) {
			return &UsageCommand{
				UI: &cli.BasicUi{
					Reader: os.Stdin,
					Writer: os.Stdout,
				},
			}, nil
		},
	}

	return commands
}
