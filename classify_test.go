package cohere

import (
	"fmt"
	"testing"
)

func TestClassify(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("ClassifySuccessMinimumFields", func(t *testing.T) {
		res, err := co.Classify(ClassifyOptions{
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
		res, err := co.Classify(ClassifyOptions{
			Model:  "large",
			Inputs: []string{"grape", "pink"},
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
		if res.Classifications[1].Prediction != "color" {
			t.Errorf("Expected: color. Receieved: %s", res.Classifications[1].Prediction)
		}
		for _, label := range []string{"fruit", "color"} {
			_, ok := res.Classifications[0].Labels[label]
			if !ok {
				fmt.Print(res.Classifications[0].Labels)
				t.Errorf("Missing confidence score for label'%s'", label)
			}
		}
	})

	t.Run("ClassifySuccessTaskDescription", func(t *testing.T) {
		res, err := co.Classify(ClassifyOptions{
			Model:  "large",
			Inputs: []string{"kiwi"},
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
		res, err := co.Classify(ClassifyOptions{
			Model:  "large",
			Inputs: []string{"pineapple"},
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

	t.Run("Classify with preset", func(t *testing.T) {
		_, err := co.Classify(ClassifyOptions{
			Preset: "SDK-TESTS-PRESET-rfa6h3",
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}
	})
}
