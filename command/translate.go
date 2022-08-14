package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/schweller/golyglot/service"
)

var _ cli.Command = (*TranslateCommand)(nil)

type TranslateCommand struct {
	UI      cli.Ui
	Command cli.Command
}

func (c *TranslateCommand) Synopsis() string {
	return "Translate a sentence or a list of sentences"
}

func (c *TranslateCommand) Help() string {
	helpText := `
		Usage: deepl translate <sentence> [options]
	`

	return strings.TrimSpace(helpText)
}

func (c *TranslateCommand) Run(args []string) int {
	var targetLang string

	flags := flag.NewFlagSet("translate", flag.ContinueOnError)

	flags.Usage = func() { c.UI.Output(c.Help()) }
	flags.StringVar(&targetLang, "target", "", "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()

	if len(args) != 1 {
		c.UI.Error("This command expects one argument: <sentence>")
		return 1
	}

	var textList []string

	textList = append(textList, args[0])
	sentence := service.Translate(textList, targetLang)
	fmt.Println(sentence)
	return 0
}
