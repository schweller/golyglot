package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/schweller/rumor"
)

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(c *rumor.Config) {
	var translateUrl string
	var translations TranslateResponse
	translateUrl = fmt.Sprintf("https://api-free.deepl.com/v2/translate?text=%s&target_lang=%s", url.QueryEscape(c.Data), c.TargetLanguage)
	if c.SourceLanguage != "" {
		translateUrl = fmt.Sprintf("https://api-free.deepl.com/v2/translate?text=%s&target_lang=%s&source_lang=%s", url.QueryEscape(c.Data), c.TargetLanguage, c.SourceLanguage)
	}
	u, err := url.Parse(translateUrl)
	if err != nil {
		log.Fatal(err)
	}

	c.Url = u
	c.Method = "POST"
	response := rumor.Execute(c)

	json.Unmarshal([]byte(response.String()), &translations)
	fmt.Println(translations.Translations[0].Text)
}
