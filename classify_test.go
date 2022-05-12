package cohere

import (
	"testing"
)

func TestClassify(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("ClassifySuccessMinimumFields", func(t *testing.T) {
		res, err := co.Classify("medium", ClassifyOptions{
			Inputs: []string{"purple"},
			Examples: []Example{
				{"apple", "fruit"}, {"banana", "fruit"}, {"watermelon", "fruit"}, {"cherry", "fruit"}, {"lemon", "fruit"},
				{"red", "color"}, {"blue", "color"}, {"blue", "color"}, {"yellow", "color"}, {"green", "color"}},
		})

		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}

		if res.Classifications[0].Prediction != "color" {
			t.Errorf("Expected: color. Receieved: %s", res.Classifications[0].Prediction)
		}
	})

	t.Run("ClassifySuccessAllFields", func(t *testing.T) {
		res, err := co.Classify("medium", ClassifyOptions{
			TaskDescription: "Classify these words as either a color or a fruit.",
			Inputs:          []string{"grape", "pink"},
			Examples: []Example{
				{"apple", "fruit"}, {"banana", "fruit"}, {"watermelon", "fruit"}, {"cherry", "fruit"}, {"lemon", "fruit"},
				{"red", "color"}, {"blue", "color"}, {"blue", "color"}, {"yellow", "color"}, {"green", "color"}},
			OutputIndicator: "This is a",
		})

		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}

		if res.Classifications[0].Prediction != "fruit" {
			t.Errorf("Expected: fruit. Receieved: %s", res.Classifications[0].Prediction)
		}
		if res.Classifications[1].Prediction != "color" {
			t.Errorf("Expected: color. Receieved: %s", res.Classifications[1].Prediction)
		}
	})

	t.Run("ClassifySuccessTaskDescription", func(t *testing.T) {
		res, err := co.Classify("medium", ClassifyOptions{
			TaskDescription: "Classify these words as a fruit or a color",
			Inputs:          []string{"kiwi"},
			Examples: []Example{
				{"apple", "fruit"}, {"banana", "fruit"}, {"watermelon", "fruit"}, {"cherry", "fruit"}, {"lemon", "fruit"},
				{"red", "color"}, {"blue", "color"}, {"blue", "color"}, {"yellow", "color"}, {"green", "color"}},
		})

		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}

		if res.Classifications[0].Prediction != "fruit" {
			t.Errorf("Expected: fruit. Receieved: %s", res.Classifications[0].Prediction)
		}
	})

	t.Run("ClassifySuccessOutputIndicator", func(t *testing.T) {
		res, err := co.Classify("medium", ClassifyOptions{
			Inputs: []string{"pineapple"},
			Examples: []Example{
				{"apple", "fruit"}, {"banana", "fruit"}, {"watermelon", "fruit"}, {"cherry", "fruit"}, {"lemon", "fruit"},
				{"red", "color"}, {"blue", "color"}, {"blue", "color"}, {"yellow", "color"}, {"green", "color"}},
			OutputIndicator: "This is a",
		})

		if err != nil {
			t.Errorf("Expected result, got error: %s", err.Error())
		}

		if res.Classifications[0].Prediction != "fruit" {
			t.Errorf("Expected: fruit. Receieved: %s", res.Classifications[0].Prediction)
		}
	})
}
