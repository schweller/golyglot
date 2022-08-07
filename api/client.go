package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type Client struct {
	config *Config
	host   string
	token  string
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
	defaultConfig := DefaultConfig()

	if c == nil {
		c = defaultConfig
	}

	client := &Client{
		config: c,
		host:   "api-free.deepl.com",
		token:  GetAuthToken(),
	}
	return client, nil
}

func (c *Client) NewRequest(method string, path string) *Request {
	req := &Request{
		Method: method,
		Path:   path,
		Host:   c.host,
		Token:  c.token,
		Params: make(map[string][]string),
	}

	return req
}

func (c *Client) MakeRequest(r *Request) (*http.Response, error) {

	httpClient := c.config.HttpClient

	req, _ := r.NewRequestHTTP()
	if os.Getenv("DEBUG") == "true" {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := httpClient.Do(req)
	if os.Getenv("DEBUG") == "true" {
		debug(httputil.DumpResponse(resp, true))
	}

	if err != nil {
		return resp, err
	}

	return resp, err
}

func (c *Client) SetClientToken(token string) {
	c.token = token
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
