package command

import (
	"fmt"

	"github.com/schweller/rumor"
	"github.com/spf13/cobra"
)

var (
	authKey    string
	sourceLang string
	targetLang string
)

func CreateCommand() *cobra.Command {
	config := rumor.Config{}

	textCommand := &cobra.Command{
		Use:     `text`,
		Args:    Validator(&config),
		PreRunE: OptionsValidator(&config),
		Run: func(cmd *cobra.Command, args []string) {
			Translate(&config)
		},
	}

	usageCommand := &cobra.Command{
		Use:     `usage`,
		PreRunE: OptionsValidator(&config),
		Run: func(cmd *cobra.Command, args []string) {
			Usage(&config)
		},
	}

	var rootCommand = &cobra.Command{Use: `rumor`}
	rootCommand.PersistentFlags().StringVarP(&authKey, "auth", "k", "", "the deepl key")
	textCommand.PersistentFlags().StringVarP(&targetLang, "to", "t", "", "which language translate to")
	textCommand.PersistentFlags().StringVarP(&sourceLang, "from", "f", "", "which language translate from")
	rootCommand.AddCommand(textCommand)
	rootCommand.AddCommand(usageCommand)

	return rootCommand
}

func Validator(c *rumor.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		headers := map[string]string{
			"Authorization": fmt.Sprintf("DeepL-Auth-Key %s", authKey),
		}
		if l := len(args); l != 1 {
			return rumor.NewErrorWithCode(2, "you must provide a single URL to be called but you provided %v", l)
		}

		c.TargetLanguage = targetLang
		c.SourceLanguage = sourceLang
		c.Headers = headers
		c.Data = args[0]

		return nil
	}
}

func OptionsValidator(c *rumor.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		headers := map[string]string{
			"Authorization": fmt.Sprintf("DeepL-Auth-Key %s", authKey),
		}
		c.Headers = headers

		return nil
	}
}
