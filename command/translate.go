package command

import (
	"fmt"

	"github.com/schweller/rumor/service"
)

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(text string, tlang string) {
	sentence := service.Translate(text, tlang)
	fmt.Println(sentence)
}
