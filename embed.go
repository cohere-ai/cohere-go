package cohere

type EmbedOptions struct {
	// An array of strings for the model to embed.
	Texts []string `json:"texts"`
}

type EmbedResponse struct {
	// An array of embeddings, where each embedding is an array of floats. The length of the embeddings
	// array will be the same as the length of the original texts array.
	Embeddings [][]float64
}
