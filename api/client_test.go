package api

import (
	"net/http"
	"testing"
)

func TestNewClientWithConfig(t *testing.T) {
	_, err := NewClient(&Config{
		HttpClient: http.DefaultClient,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewClientConfigNil(t *testing.T) {
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal("expected a non-nil client")
	}
}
