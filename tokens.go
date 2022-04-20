package cohere

type TokenizeOptions struct {
	// The string to be tokenized
	Text string `json:"text"`
}

type TokenizeResponse struct {
	// The tokens.
	Tokens   []string `json:"tokens"`
	TokenIDs []int64  `json:"token_ids"`
}
