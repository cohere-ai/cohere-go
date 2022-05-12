package cohere

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("Embed", func(t *testing.T) {
		texts := []string{"hello", "goodbye"}

		_, err := co.Embed("small", EmbedOptions{
			Texts:    texts,
			Truncate: TruncateNone,
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}
