package rumor

import (
	"fmt"
	"log"
	"net/url"

	"github.com/imroc/req/v3"
)

type Client struct {
	client *req.Client
}

func NewClient() *Client {
	return &Client{
		client: req.C(),
	}
}

type Config struct {
	Headers        map[string]string
	TargetLanguage string
	SourceLanguage string
	Method         string
	Url            *url.URL
	Endpoint       string
	Data           string
}

func (r *Client) Execute(c *Config) *req.Response {
	r.client.SetCommonHeaders(c.Headers)

	finalUrl := buildUrl(c.Endpoint)
	response, err := r.client.
		R().
		Send(c.Method, finalUrl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func buildUrl(r string) string {
	finalUrl := ""

	finalUrl = fmt.Sprintf("https://api-free.deepl.com/v2%s", r)

	return finalUrl
}
