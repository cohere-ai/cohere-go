package cohere

import (
	"errors"
	"os"
	"reflect"
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
		co := CreateClient("")
		_, err := co.Generate("baseline-shrimp", GenerateOptions{
			Prompt:      "",
			MaxTokens:   10,
			Temperature: 0.75,
		})
		if err == nil {
			t.Errorf("expected error, got nil")
		} else if !errors.Is(err, &APIError{}) {
			t.Errorf("expected ApiError, got %s", reflect.TypeOf(err))
		}
	})
}

func TestGenerate(t *testing.T) {
	co := CreateClient(apiKey)

	t.Run("Generate basic", func(t *testing.T) {
		_, err := co.Generate("baseline-shark", GenerateOptions{
			Prompt:      "Hello my name is",
			MaxTokens:   10,
			Temperature: 0.75,
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestSimilarity(t *testing.T) {
	co := CreateClient(apiKey)

	t.Run("Similarity", func(t *testing.T) {
		_, err := co.Similarity("baseline-shrimp", SimilarityOptions{
			Anchor:  "hi how are you doing today?",
			Targets: []string{"greeting", "request for assistance", "asking a question"},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestChooseBest(t *testing.T) {
	co := CreateClient(apiKey)

	t.Run("ChooseBest", func(t *testing.T) {
		_, err := co.ChooseBest("baseline-otter", ChooseBestOptions{
			Query:   "Carol picked up a book and walked to the kitchen. She set it down, picked up her glasses and left. This is in the kitchen now: ",
			Options: []string{"book", "glasses", "dog"},
			Mode:    AppendOption,
		})

		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestEmbed(t *testing.T) {
	co := CreateClient(apiKey)

	t.Run("Embed", func(t *testing.T) {
		texts := []string{"hello", "goodbye"}

		_, err := co.Embed("baseline-shrimp", texts)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestLikelihood(t *testing.T) {
	co := CreateClient(apiKey)

	t.Run("Likelihood", func(t *testing.T) {
		text := "so I crept up the basement stairs and BOOOO!"

		_, err := co.Likelihood("baseline-orca", text)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}
