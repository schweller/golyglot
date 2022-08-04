package service

import (
	"encoding/json"

	"github.com/schweller/rumor"
)

type Foo struct {
}

type UsageResponse struct {
	Count int `json:"character_count"`
	Limit int `json:"character_limit"`
}

func ShowUsage(c *rumor.Config) (int, int) {
	var usageInfo UsageResponse

	client := rumor.NewClient()

	c.Method = "GET"
	c.Endpoint = "/usage"
	response := client.Execute(c)

	json.Unmarshal([]byte(response.String()), &usageInfo)
	return usageInfo.Count, usageInfo.Limit
}
