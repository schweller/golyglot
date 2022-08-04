package rumor

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	key := new(string)
	config := Config{}

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
	rootCommand.PersistentFlags().StringVarP(key, "auth", "k", "", "the deepl key")
	rootCommand.AddCommand(textCommand)
	rootCommand.AddCommand(usageCommand)

	return rootCommand
}

func Validator(c *Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if l := len(args); l != 1 {
			return newErrorWithCode(2, "you must provide a single URL to be called but you provided %v", l)
		}

		c.Data = args[0]

		return nil
	}
}

func OptionsValidator(c *Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		headers := map[string]string{
			"Authorization": "DeepL-Auth-Key",
		}
		c.Headers = headers

		return nil
	}
}
