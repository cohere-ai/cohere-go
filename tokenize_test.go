package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	t.Run("TokenizeSuccess", func(t *testing.T) {
		text := "hello world"

		res, err := Tokenize(TokenizeOptions{
			Text: text,
		})
		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}
		expectedTokens := []int64{33555, 1114}
		expectedTokenStrings := []string{"hello", " world"}
		assert.Equal(t, expectedTokens, res.Tokens)
		assert.Equal(t, expectedTokenStrings, res.TokenStrings)
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
