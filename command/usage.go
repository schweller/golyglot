package command

import (
	"fmt"

	"github.com/schweller/rumor"
	"github.com/schweller/rumor/service"
)

type UsageResponse struct {
	Count int `json:"character_count"`
	Limit int `json:"character_limit"`
}

func Usage(c *rumor.Config) {
	count, limit := service.ShowUsage(c)
	fmt.Println(fmt.Sprintf("You already used %v characters out of your limit %v", count, limit))
}
