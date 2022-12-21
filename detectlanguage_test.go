package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectLanguage(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("DetectLanguage success", func(t *testing.T) {
		res, err := co.DetectLanguage(DetectLanguageOptions{
			Texts: []string{"this text is in english", "Этот текст на Русском языке."},
		})
		assert.Nil(t, err)
		assert.Equal(t, res.Results[0].LanguageCode, "en")
		assert.Equal(t, res.Results[0].LanguageName, "English")
		assert.Positive(t, res.Results[0].Confidence)
		assert.Equal(t, res.Results[1].LanguageCode, "ru")
		assert.Equal(t, res.Results[1].LanguageName, "Russian")
		assert.Positive(t, res.Results[1].Confidence)
	})
}
