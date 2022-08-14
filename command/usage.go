package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/schweller/golyglot/service"
)

type UsageCommand struct {
	UI      cli.Ui
	Command cli.Command
}

func (c *UsageCommand) Synopsis() string {
	return "Translate a sentence or a list of sentences"
}

func (c *UsageCommand) Help() string {
	helpText := `
		Usage: deepl translate <sentence> [options]
	`

	return strings.TrimSpace(helpText)
}

func (c *UsageCommand) Run(args []string) int {
	flags := flag.NewFlagSet("usage", flag.ContinueOnError)

	flags.Usage = func() { c.UI.Output(c.Help()) }

	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()

	if len(args) != 0 {
		c.UI.Error("This command takes no argument!")
		return 1
	}

	count, limit := service.GetUsage()
	fmt.Println(fmt.Sprintf("You already used %v characters out of your limit %v", count, limit))
	return 0
}
