package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Client struct {
	config *Config
	host   string
}

type Config struct {
	Type       string
	HttpClient *http.Client
}

func DefaultConfig() *Config {
	return &Config{
		HttpClient: getCleanClient(),
	}
}

func getCleanClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{},
	}
}

func NewClient(c *Config) (*Client, error) {
	// config := DefaultConfig()
	client := &Client{
		config: c,
		host:   "api-free.deepl.com",
	}
	return client, nil
}

type Request struct {
	Method string
	URL    *url.URL
	Path   string
	Host   string
	Params url.Values
	Body   io.Reader
}

type Response struct {
	*http.Response
}

func NewRequest(r *Request) *http.Request {
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("DeepL-Auth-Key %s", GetAuthToken())},
	}

	rawQuery := ""

	if r.Params != nil {
		rawQuery = r.Params.Encode()
	}

	req := &http.Request{
		Method: r.Method,
		URL: &url.URL{
			Host:     r.Host,
			Scheme:   "https",
			Path:     r.Path,
			RawQuery: rawQuery,
		},
		Header: headers,
	}
	return req
}

func (c *Client) MakeRequest(r *Request) (*Response, error) {

	httpClient := c.config.HttpClient

	req := NewRequest(r)

	debug(httputil.DumpRequestOut(req, true))
	resp, err := httpClient.Do(req)

	var result *Response
	if resp != nil {
		result = &Response{Response: resp}
	}

	if err != nil {
		return result, err
	}

	return result, err
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
