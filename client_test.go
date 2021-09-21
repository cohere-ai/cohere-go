package cohere

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

var API_KEY = os.Getenv("API_KEY")

func init() {
	if API_KEY == "" {
		panic("API_KEY is not set")
	}
}

func TestErrors(t *testing.T) {
	t.Run("Invalid api key", func(t *testing.T) {
		co := CreateClient("")
		_, err := co.Generate(Shrimp, "", 10, 0.75)
		if err == nil {
			t.Errorf("expected error, got nil")
		} else if !errors.Is(err, &ApiError{}) {
			t.Errorf("expected ApiError, got %s", reflect.TypeOf(err))
		}
	})
}

func TestGenerate(t *testing.T) {
	co := CreateClient(API_KEY)

	t.Run("Generate basic", func(t *testing.T) {
		_, err := co.Generate(Orca, "Hello my name is", 10, 0.75)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})

	t.Run("Generate advanced", func(*testing.T) {
		maxTokens := uint(10)

		_, err := co.GenerateAdvanced(Seal, GenerateOptions{
			Prompt:            "What is your",
			MaxTokens:         maxTokens,
			Temperature:       1,
			K:                 5,
			P:                 0,
			StopSequences:     []string{".", "?"},
			ReturnLikelihoods: ALL,
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestSimilarity(t *testing.T) {
	co := CreateClient(API_KEY)

	t.Run("Similarity", func(t *testing.T) {
		anchor := "hi how are you doing today?"
		targets := []string{"greeting", "request for assistance", "asking a question"}

		_, err := co.Similarity(Seal, anchor, targets)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestChooseBest(t *testing.T) {
	co := CreateClient(API_KEY)

	t.Run("ChooseBest", func(t *testing.T) {
		query := "Carol picked up a book and walked to the kitchen. She set it down, picked up her glasses and left. This is in the kitchen now: "
		options := []string{"book", "glasses", "dog"}
		mode := APPEND_OPTION

		_, err := co.ChooseBest(Otter, query, options, mode)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestEmbed(t *testing.T) {
	co := CreateClient(API_KEY)

	t.Run("Embed", func(t *testing.T) {
		texts := []string{"hello", "goodbye"}

		_, err := co.Embed(Seal, texts)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestLikelihood(t *testing.T) {
	co := CreateClient(API_KEY)

	t.Run("Likelihood", func(t *testing.T) {
		text := "so I crept up the basement stairs and BOOOO!"

		_, err := co.Likelihood(Orca, text)
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}
