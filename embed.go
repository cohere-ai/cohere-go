package cohere

// return likelihoods
const (
	TruncateNone  = "NONE"
	TruncateLeft  = "LEFT"
	TruncateRight = "RIGHT"
)

type EmbedOptions struct {
	// An optional string representing the model you'd like to use.
	Model string `json:"model,omitempty"`
	// An array of strings for the model to embed.
	Texts    []string `json:"texts"`
	Truncate string   `json:"truncate"`
}

type EmbedResponse struct {
	// An array of embeddings, where each embedding is an array of floats. The length of the embeddings
	// array will be the same as the length of the original texts array.
	Embeddings [][]float64

	// Metadata about the API version
	Meta *MetaResponse `json:"meta,omitempty"`
}
