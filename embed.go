package cohere

type EmbedOptions struct {
	// An optional string representing the model you'd like to use.
	Model string `json:"model,omitempty"`

	// An array of strings for the model to embed.
	Texts []string `json:"texts"`

	// Accepts one of TruncateNone, TruncateLeft and TruncateRight. Determines how inputs over the maximum context length will be handled.
	// Passing TruncateLeft will discard the left of the input and TruncateRight will discard the right side of the input.
	// TruncateNone will return an error if the input is too long.
	Truncate string `json:"truncate,omitempty"`
}

type EmbedResponse struct {
	// An array of embeddings, where each embedding is an array of floats. The length of the embeddings
	// array will be the same as the length of the original texts array.
	Embeddings [][]float64
}
