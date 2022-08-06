package api

import (
	"log"
	"net/url"
)

type TranslateApi struct {
	client *Client
}

type Translations struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string
}

type TranslateApiResponse struct {
	Translations []Translations
}

type Translate struct {
	client *Client
}

func (c *Client) Translate() *UsageApi {
	return &UsageApi{
		client: c,
	}
}

func (api *UsageApi) PostTranslation(text string, target_lang string) (*TranslateApiResponse, error) {
	params := url.Values{}

	params.Add("text", text)
	params.Add("target_lang", target_lang)

	req := &Request{
		Method: "POST",
		Path:   "/v2/translate",
		Host:   api.client.host,
		Params: params,
	}

	response, err := api.client.MakeRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	var result TranslateApiResponse
	if err := DecodeJSON(response.Body, &result); err != nil {
		return nil, err
	}

	return &result, err
}
