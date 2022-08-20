package cohere

type TokenizeOptions struct {
	// The string to be tokenized
	Text string `json:"text"`
}

type TokenizeResponse struct {
	// The tokens
	Tokens []int64 `json:"tokens"`
	// String representations of the tokens
	TokenStrings []string `json:"tokenStrings"`
}
