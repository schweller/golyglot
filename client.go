package rumor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/imroc/req/v3"
)

type Config struct {
	Headers            map[string]string
	Method             string
	Url                *url.URL
	Data               string
	ResponseBodyOutput io.Writer
	ControlOutput      io.Writer
}

func execute(c *Config) *req.Response {
	client := req.C().
		SetCommonHeaders(c.Headers)

	response, err := client.
		R().
		Send(c.Method, c.Url.String())

	if err != nil {
		log.Fatal(err)
	}

	return response
}

type UsageResponse struct {
	CharacterCount int `json:"character_count"`
	CharacterLimit int `json:"character_limit"`
}

func Usage(c *Config) {
	var usageInfo UsageResponse

	u, err := url.Parse("https://api-free.deepl.com/v2/usage")
	if err != nil {
		log.Fatal(err)
	}
	c.Url = u
	c.Method = "GET"
	response := execute(c)

	json.Unmarshal([]byte(response.String()), &usageInfo)
	fmt.Println(fmt.Sprintf("You already used %v characters out of your limit %v", usageInfo.CharacterCount, usageInfo.CharacterLimit))
}

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateResponse struct {
	Translations []Translations
}

func Translate(c *Config) {
	var translations TranslateResponse
	u, err := url.Parse(fmt.Sprintf("https://api-free.deepl.com/v2/translate?text=%s&source_lang=EN&target_lang=DE", url.QueryEscape(c.Data)))
	if err != nil {
		log.Fatal(err)
	}

	c.Url = u
	c.Method = "POST"
	response := execute(c)

	json.Unmarshal([]byte(response.String()), &translations)
	fmt.Println(translations.Translations[0].Text)
}
