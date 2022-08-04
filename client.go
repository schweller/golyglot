package rumor

import (
	"io"
	"log"
	"net/url"

	"github.com/imroc/req/v3"
)

type Config struct {
	Headers            map[string]string
	TargetLanguage     string
	SourceLanguage     string
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

func Execute(c *Config) *req.Response {
	return execute(c)
}
