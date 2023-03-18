package cohere

type EmbedOptions struct {
	// An optional string representing the model you'd like to use.
	Model string `json:"model,omitempty"`

	// An array of strings for the model to embed.
	Texts []string `json:"texts"`

	// Accepts one of TruncateNone, TruncateStart and TruncateEnd. Determines how inputs over the maximum context length will be handled.
	// Passing TruncateStart will discard the start of the input and TruncateEnd will discard the end of the input.
	// Defaults to TruncateNone, which will return an error if the input is too long.
	Truncate string `json:"truncate,omitempty"`
}

type EmbedResponse struct {
	// An array of embeddings, where each embedding is an array of floats. The length of the embeddings
	// array will be the same as the length of the original texts array.
	Embeddings [][]float64

	// Metadata about the API version
	Meta *MetaResponse `json:"meta,omitempty"`
}
