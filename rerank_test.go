package cohere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReRank(t *testing.T) {
	co, err := CreateClient(apiKey)
	assert.NoError(t, err, "failed to instantiate cohere client")

	t.Run("ReRank", func(t *testing.T) {
		return_documents := true
		query := "What is the capital of the United States?"
		documents := []interface{}{
			"Carson City is the capital city of the American state of Nevada.",
			"The Commonwealth of the Northern Mariana Islands is a group of islands in the Pacific Ocean. Its capital is Saipan.",
			"Washington, D.C. (also known as simply Washington or D.C., and officially as the District of Columbia) is the capital of the United States. It is a federal district.",
			"Capital punishment (the death penalty) has existed in the United States since beforethe United States was a country. As of 2017, capital punishment is legal in 30 of the 50 states.",
		}

		res, err := co.Rerank(RerankOptions{
			Query:           query,
			Documents:       documents,
			ReturnDocuments: return_documents,
		})

		assert.NoError(t, err, "expected successful reranking")
		assert.NotEmpty(t, res.ID)
		assert.NotEmpty(t, res.Results)
	})
}
