package command

import (
	"os"

	"github.com/schweller/rumor"
	"github.com/spf13/cobra"
)

var (
	authKey    string
	sourceLang string
	targetLang string
)

func CreateRootCommand() *cobra.Command {
	config := rumor.Config{}

	textCommand := &cobra.Command{
		Use:  `text`,
		Args: Validator(&config),
		Run: func(cmd *cobra.Command, args []string) {
			Translate(args[0], targetLang)
		},
	}

	usageCommand := &cobra.Command{
		Use: `usage`,
		Run: func(cmd *cobra.Command, args []string) {
			Usage(&config)
		},
	}

	var rootCommand = &cobra.Command{
		Use:               `rumor`,
		PersistentPreRunE: InitialValidator(&config),
	}
	textCommand.PersistentFlags().StringVarP(&targetLang, "to", "t", "", "which language translate to")
	textCommand.PersistentFlags().StringVarP(&sourceLang, "from", "f", "", "which language translate from")
	rootCommand.AddCommand(textCommand)
	rootCommand.AddCommand(usageCommand)

	return rootCommand
}

func Validator(c *rumor.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if l := len(args); l != 1 {
			return rumor.NewErrorWithCode(2, "you must provide the text")
		}

		c.TargetLanguage = targetLang
		c.SourceLanguage = sourceLang
		c.Data = args[0]

		return nil
	}
}

func InitialValidator(c *rumor.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		key := getAuthToken()

		if key == "" {
			return rumor.NewErrorWithCode(2, "You must set DEEPL_AUTH_TOKEN environment variable")
		}

		return nil
	}
}

func getAuthToken() string {
	return os.Getenv("DEEPL_AUTH_TOKEN")
}
