package api

import (
	"log"
)

type UsageApi struct {
	client *Client
}

type UsageApiResponse struct {
	Count int `json:"character_count"`
	Limit int `json:"character_limit"`
}

type Usage struct {
	client *Client
}

func (c *Client) Usage() *UsageApi {
	return &UsageApi{
		client: c,
	}
}

func (api *UsageApi) Get() (*UsageApiResponse, error) {
	req := &Request{
		Method: "GET",
		Path:   "/v2/usage",
		Host:   api.client.host,
	}

	response, err := api.client.MakeRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	var result UsageApiResponse
	if err := DecodeJSON(response.Body, &result); err != nil {
		return nil, err
	}

	return &result, err
}
