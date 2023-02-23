package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			MaxTokens:   Uint(10),
			Temperature: Float64(0.75),
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
			MaxTokens:      Uint(10),
			Temperature:    Float64(0.75),
			NumGenerations: Int(num),
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
			MaxTokens:         Uint(10),
			Temperature:       Float64(0.75),
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
			MaxTokens:         Uint(10),
			Temperature:       Float64(0.75),
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
			MaxTokens:         Uint(10),
			Temperature:       Float64(0.75),
			ReturnLikelihoods: "NONE",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
		if res.Generations[0].Likelihood != nil {
			t.Errorf("expected nil, got %p", res.Generations[0].Likelihood)
		}
	})

	t.Run("Generate with preset", func(t *testing.T) {
		_, err := co.Generate(GenerateOptions{
			Preset: "SDK-TESTS-PRESET-cq2r57",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})

	t.Run("Generate logit bias", func(t *testing.T) {
		_, err := co.Generate(GenerateOptions{
			Model:       "medium",
			Prompt:      "Hello my name is",
			MaxTokens:   Uint(10),
			Temperature: Float64(0.75),
			LogitBias:   map[int]float32{11: -5, 33: 7.5},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}

func TestStream(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("Stream", func(t *testing.T) {
		ch := co.Stream(GenerateOptions{
			Model:       "xlarge",
			Prompt:      "Hello my name is",
			MaxTokens:   Uint(100),
			Temperature: Float64(0.9),
		})

		for res := range ch {
			require.NoError(t, res.Err)
			assert.Equal(t, res.Token.Index, 0)
			assert.NotEmpty(t, res.Token.Token)
			assert.NotZero(t, res.Token.Likelihood)
		}
	})

	t.Run("Stream multiple generations", func(t *testing.T) {
		numGens := 5
		ch := co.Stream(GenerateOptions{
			Model:          "xlarge",
			NumGenerations: Int(numGens),
			Prompt:         "Hello my name is",
			MaxTokens:      Uint(10),
			Temperature:    Float64(0.9),
		})

		seen := make(map[int]struct{})
		for res := range ch {
			require.NoError(t, res.Err)
			seen[res.Token.Index] = struct{}{}
			assert.NotEmpty(t, res.Token.Token)
			assert.NotZero(t, res.Token.Likelihood)
		}
		for i := 0; i < numGens; i++ {
			_, ok := seen[i]
			assert.True(t, ok, "missing generation with index '%d'", i)
		}
	})

	t.Run("Streaming error", func(t *testing.T) {
		ch := co.Stream(GenerateOptions{
			Model:       "non-existent-model",
			Prompt:      "Hello my name is",
			MaxTokens:   Uint(100),
			Temperature: Float64(0.9),
		})

		for res := range ch {
			require.Error(t, res.Err)
		}
	})
}
