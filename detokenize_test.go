package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetokenize(t *testing.T) {
	t.Run("DetokenizeSuccess", func(t *testing.T) {
		tokens := []int64{33555, 1114}

		res, err := Detokenize(DetokenizeOptions{
			Tokens: tokens,
		})
		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}
		expectedText := "hello world"
		assert.Equal(t, expectedText, res.Text)
	})

	t.Run("DetokenizeEmptyList", func(t *testing.T) {
		res, err := Detokenize(DetokenizeOptions{
			Tokens: []int64{},
		})
		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}
		expected := ""
		if len(res.Text) != 0 {
			t.Errorf("Detokenization failed. Expected: %v, Output: %v", res.Text, expected)
		}
	})
}
