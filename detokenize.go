package cohere

type DetokenizeOptions struct {
	// The tokens to be detokenized
	Tokens []int64 `json:"tokens"`
}

type DetokenizeResponse struct {
	// The text represention of the tokens
	Text string `json:"text"`

	// Metadata about the API version
	Meta *MetaResponse `json:"meta,omitempty"`
}
