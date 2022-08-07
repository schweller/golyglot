package command

import (
	"fmt"

	"github.com/schweller/golyglot/service"
)

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(text string, tlang string) {
	var textList []string

	textList = append(textList, text)

	sentence := service.Translate(textList, tlang)
	fmt.Println(sentence)
}
