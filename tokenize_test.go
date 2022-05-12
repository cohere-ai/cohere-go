package cohere

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Run("TokenizeSuccess", func(t *testing.T) {
		text := "tokenize me!"

		_, err := Tokenize(TokenizeOptions{
			Text: text,
		})
		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}
	})

	t.Run("TokenizeEmptyText", func(t *testing.T) {
		text := ""

		res, err := Tokenize(TokenizeOptions{
			Text: text,
		})
		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}
		expected := []int64{}
		if len(res.Tokens) != 0 {
			t.Errorf("Tokenization failed. Expected: %v, Output: %v", res.Tokens, expected)
		}
	})
}
