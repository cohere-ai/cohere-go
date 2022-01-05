package cohere

type TokenizeOptions struct {
	//The string to be tokenized
	Text string `json:"text"`
}

type TokenizeResponse struct {
	//The tokens.
	Tokens []uint `json:"tokens"`
}
