package service

import (
	"encoding/json"

	"github.com/schweller/rumor"
)

type UsageResponse struct {
	Count int `json:"character_count"`
	Limit int `json:"character_limit"`
}

func ShowUsage(c *rumor.Config) (int, int) {
	usageInfo := UsageResponse{}

	client := rumor.NewClient()

	response := client.Get(c, "/usage")

	json.Unmarshal([]byte(response.String()), &usageInfo)
	return usageInfo.Count, usageInfo.Limit
}
