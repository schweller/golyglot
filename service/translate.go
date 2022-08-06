package service

import (
	"log"

	"github.com/schweller/rumor/api"
)

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(text string, tlang string) string {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize client: %v", err)
	}

	resp, bar := client.Translate().GetTranslation(text, tlang)

	if bar != nil {
		log.Fatal(bar)
	}

	return resp.Translations[0].Text
}
