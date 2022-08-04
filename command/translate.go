package command

import (
	"fmt"

	"github.com/schweller/rumor"
	"github.com/schweller/rumor/service"
)

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(c *rumor.Config) {
	text := service.Translate(c)
	fmt.Println(text)
}
