package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	co, err := CreateClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	t.Run("Extraction with single text", func(t *testing.T) {
		res, err := co.Extract("medium", ExtractOptions{
			Texts: []string{"Jim just came back from soccer practice."},
			Examples: []ExtractExample{
				{
					Text: "hello my name is John, and I like to play ping pong",
					Entities: []ExtractEntity{
						{Type: "name", Value: "John"},
						{Type: "game", Value: "ping pong"},
					},
				},
				{
					Text: "Karen is the best tennis player in her grade",
					Entities: []ExtractEntity{
						{Type: "name", Value: "Karen"},
						{Type: "game", Value: "tennis"},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}

		assert.Equal(t, 1, len(res.Extractions), "Expected 1 extraction")
		assert.Equal(t, "Jim just came back from soccer practice.", res.Extractions[0].Text, "Expected text to be the same")
		assert.NotNil(t, res.Extractions[0].ID, "Extraction expected to have an id")
		assert.NotNil(t, res.Extractions[0].Entities, "Extraction expected to have entities")
		assert.Contains(t, res.Extractions[0].Entities, ExtractEntity{Type: "name", Value: "Jim"}, "Expected extraction to contain name")
		assert.Contains(t, res.Extractions[0].Entities, ExtractEntity{Type: "game", Value: "soccer"}, "Expected extraction to contain game")
	})

	t.Run("Extraction with multiple text inputs", func(t *testing.T) {
		res, err := co.Extract("medium", ExtractOptions{
			Texts: []string{"Jim just came back from soccer practice.", "Who knew that Scott would be so good at chess??"},
			Examples: []ExtractExample{
				{
					Text: "hello my name is John, and I like to play ping pong",
					Entities: []ExtractEntity{
						{Type: "name", Value: "John"},
						{Type: "game", Value: "ping pong"},
					},
				},
				{
					Text: "Karen is the best tennis player in her grade",
					Entities: []ExtractEntity{
						{Type: "name", Value: "Karen"},
						{Type: "game", Value: "tennis"},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}

		assert.Equal(t, 2, len(res.Extractions), "Expected 2 extractions")

		assert.Equal(t, "Jim just came back from soccer practice.", res.Extractions[0].Text, "Expected text to be the same")
		assert.NotNil(t, res.Extractions[0].ID, "Extraction expected to have an id")
		assert.NotNil(t, res.Extractions[0].Entities, "Extraction expected to have entities")
		assert.Contains(t, res.Extractions[0].Entities, ExtractEntity{Type: "name", Value: "Jim"}, "Expected extraction to contain name")
		assert.Contains(t, res.Extractions[0].Entities, ExtractEntity{Type: "game", Value: "soccer"}, "Expected extraction to contain game")

		assert.Equal(t, "Who knew that Scott would be so good at chess??", res.Extractions[1].Text, "Expected text to be the same")
		assert.NotNil(t, res.Extractions[1].ID, "Extraction expected to have an id")
		assert.NotNil(t, res.Extractions[1].Entities, "Extraction expected to have entities")
		assert.Contains(t, res.Extractions[1].Entities, ExtractEntity{Type: "name", Value: "Scott"}, "Expected extraction to contain name")
		assert.Contains(t, res.Extractions[1].Entities, ExtractEntity{Type: "game", Value: "chess"}, "Expected extraction to contain game")
	})

	t.Run("Run extraction on text with partial entity matching", func(t *testing.T) {
		res, err := co.Extract("medium", ExtractOptions{
			Texts: []string{"Jared just returned."},
			Examples: []ExtractExample{
				{
					Text: "hello my name is John, and I like to play ping pong",
					Entities: []ExtractEntity{
						{Type: "name", Value: "John"},
						{Type: "game", Value: "ping pong"},
					},
				},
				{
					Text: "Karen is the best tennis player in her grade",
					Entities: []ExtractEntity{
						{Type: "name", Value: "Karen"},
						{Type: "game", Value: "tennis"},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}

		assert.Len(t, res.Extractions, 1, "Expected 1 extraction")
		assert.Equal(t, "Jared just returned.", res.Extractions[0].Text, "Expected text to be the same")
		assert.NotNil(t, res.Extractions[0].ID, "Extraction expected to have an id")
		assert.NotNil(t, res.Extractions[0].Entities, "Extraction expected to have entities")
		assert.Len(t, res.Extractions[0].Entities, 1, "Expected extraction to contain 1 entity")
		assert.Contains(t, res.Extractions[0].Entities, ExtractEntity{Type: "name", Value: "Jared"}, "Expected extraction to contain name")
	})

	t.Run("Run extraction on text with no matched entities", func(t *testing.T) {
		res, err := co.Extract("medium", ExtractOptions{
			Texts: []string{"Shirt size XXL."},
			Examples: []ExtractExample{
				{
					Text: "hello my name is John, and I like to play ping pong",
					Entities: []ExtractEntity{
						{Type: "name", Value: "John"},
						{Type: "game", Value: "ping pong"},
					},
				},
				{
					Text: "Karen is the best tennis player in her grade",
					Entities: []ExtractEntity{
						{Type: "name", Value: "Karen"},
						{Type: "game", Value: "tennis"},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("expected result, got error: %s", err.Error())
		}

		assert.Len(t, res.Extractions, 1, "Expected 1 extraction")
		assert.Equal(t, "Shirt size XXL.", res.Extractions[0].Text, "Expected text to be the same")
		assert.NotNil(t, res.Extractions[0].ID, "Extraction expected to have an id")
		assert.Empty(t, res.Extractions[0].Entities, "Extraction expected to have no entities")
	})
}
