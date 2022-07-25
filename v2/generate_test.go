package cohere

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("Generate basic", func(t *testing.T) {
		_, err := co.Generate(GenerateOptions{
			Model:       "medium",
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
		res, err := co.Generate(GenerateOptions{
			Model:          "medium",
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
		res, err := co.Generate(GenerateOptions{
			Model:             "medium",
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
		res, err := co.Generate(GenerateOptions{
			Model:             "medium",
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
		res, err := co.Generate(GenerateOptions{
			Model:             "medium",
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
