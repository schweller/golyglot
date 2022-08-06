package api

import (
	"encoding/json"
	"io"
	"os"
)

func DecodeJSON(reader io.Reader, output interface{}) error {
	decoder := json.NewDecoder(reader)

	return decoder.Decode(output)
}

func GetAuthToken() string {
	return os.Getenv("DEEPL_AUTH_TOKEN")
}
