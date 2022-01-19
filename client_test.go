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

func TestGenerate(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("Generate basic", func(t *testing.T) {
		_, err := co.Generate("medium", GenerateOptions{
			Prompt:      "Hello my name is",
			MaxTokens:   10,
			Temperature: 0.75,
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})

	t.Run("Generate multi", func(t *testing.T) {
		num := 4
		res, err := co.Generate("medium", GenerateOptions{
			Prompt:         "What is your",
			MaxTokens:      10,
			Temperature:    0.75,
			NumGenerations: num,
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		} else if len(res.Generations) != num {
			t.Errorf("expected %d gnerations, got %d", num, len(res.Generations))
		}
	})

	t.Run("Generate likelihood with generation", func(t *testing.T) {
		res, err := co.Generate("medium", GenerateOptions{
			Prompt:            "Hello my name is",
			MaxTokens:         10,
			Temperature:       0.75,
			ReturnLikelihoods: "GENERATION",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
		if res.Generations[0].Likelihood == nil {
			t.Errorf("expected likelihood")
		}
	})

	t.Run("Generate likelihood with all", func(t *testing.T) {
		res, err := co.Generate("medium", GenerateOptions{
			Prompt:            "Hello my name is",
			MaxTokens:         10,
			Temperature:       0.75,
			ReturnLikelihoods: "ALL",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
		if res.Generations[0].Likelihood == nil {
			t.Errorf("expected likelihood")
		}
	})

	t.Run("Generate likelihood with none", func(t *testing.T) {
		res, err := co.Generate("medium", GenerateOptions{
			Prompt:            "Hello my name is",
			MaxTokens:         10,
			Temperature:       0.75,
			ReturnLikelihoods: "NONE",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
		if res.Generations[0].Likelihood != nil {
			t.Errorf("expected nil, got %p", res.Generations[0].Likelihood)
		}
	})
}

func TestChooseBest(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("ChooseBest", func(t *testing.T) {
		_, err := co.ChooseBest("small", ChooseBestOptions{
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

func TestTokenize(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("TokenizeSuccess", func(t *testing.T) {
		text := "tokenize me!"

		_, err := co.Tokenize("large", TokenizeOptions{
			Text: text,
		})
		if err != nil {
			t.Errorf("expected result, go error: %s", err.Error())
		}
	})

	t.Run("TokenizeEmptyText", func(t *testing.T) {
		text := ""

		_, err := co.Tokenize("large", TokenizeOptions{
			Text: text,
		})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
