package service

import (
	"log"

	"github.com/schweller/golyglot/api"
)

func GetUsage() (int, int) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)

	if err != nil {
		log.Fatalf("unable to initialize client: %v", err)
	}

	resp, bar := client.Usage().Get()

	if bar != nil {
		log.Fatal(bar)
	}

	return resp.Count, resp.Limit
}
