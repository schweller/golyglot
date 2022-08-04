package rumor

import (
	"fmt"
	"log"
	"os"

	"github.com/imroc/req/v3"
)

type Client struct {
	client *req.Client
	key    string
}

func NewClient() *Client {
	return &Client{
		client: req.C(),
		key:    getAuthToken(),
	}
}

type Config struct {
	Headers        map[string]string
	TargetLanguage string
	SourceLanguage string
	Data           string
}

func (rumor *Client) Get(c *Config, path string) *req.Response {
	return rumor.Execute(c, "GET", path)
}

func (rumor *Client) Post(c *Config, path string) *req.Response {
	return rumor.Execute(c, "POST", path)
}

func (rumor *Client) Execute(c *Config, method string, path string) *req.Response {
	addHeadersToRequest(rumor)

	finalUrl := buildUrl(path)
	response, err := rumor.client.
		R().
		Send(method, finalUrl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func addHeadersToRequest(rumor *Client) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("DeepL-Auth-Key %s", rumor.key),
	}

	rumor.client.SetCommonHeaders(headers)
}

func getAuthToken() string {
	return os.Getenv("DEEPL_AUTH_TOKEN")
}

func buildUrl(r string) string {
	finalUrl := ""

	finalUrl = fmt.Sprintf("https://api-free.deepl.com/v2%s", r)

	return finalUrl
}
