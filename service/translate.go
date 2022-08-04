package service

import (
	"encoding/json"
	"fmt"
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

func Translate(c *rumor.Config) string {
	client := rumor.NewClient()
	translateUrl := ""
	translations := &TranslateResponse{}

	translateUrl = fmt.Sprintf("/translate?text=%s&target_lang=%s", url.QueryEscape(c.Data), c.TargetLanguage)
	if c.SourceLanguage != "" {
		translateUrl = fmt.Sprintf("/translate?text=%s&target_lang=%s&source_lang=%s", url.QueryEscape(c.Data), c.TargetLanguage, c.SourceLanguage)
	}

	response := client.Post(c, translateUrl)

	json.Unmarshal([]byte(response.String()), &translations)

	return translations.Translations[0].Text
}
