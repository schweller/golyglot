package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/schweller/rumor"
)

type UsageResponse struct {
	CharacterCount int `json:"character_count"`
	CharacterLimit int `json:"character_limit"`
}

func Usage(c *rumor.Config) {
	var usageInfo UsageResponse

	u, err := url.Parse("https://api-free.deepl.com/v2/usage")
	if err != nil {
		log.Fatal(err)
	}
	c.Url = u
	c.Method = "GET"
	response := rumor.Execute(c)

	json.Unmarshal([]byte(response.String()), &usageInfo)
	fmt.Println(fmt.Sprintf("You already used %v characters out of your limit %v", usageInfo.CharacterCount, usageInfo.CharacterLimit))
}
