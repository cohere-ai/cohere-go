package cohere

import (
	"os"
	"testing"
)

var apiKey = os.Getenv("API_KEY")

func init() {
	if apiKey == "" {
		panic("api key is not set")
	}
}

func TestErrors(t *testing.T) {
	t.Run("Invalid api key", func(t *testing.T) {
		co, err := CreateClient("")
		if co != nil {
			t.Error("expected nil client, got client")
		}
		if err == nil {
			t.Error("expected error, got nil")
		} else if err.Error() != "invalid api key" {
			t.Errorf("expected invalid api key, got %s", err.Error())
		}
	})
}
