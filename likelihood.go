package cohere

type TokenLikelihood struct {
	// The token.
	Token string `json:"token"`

	// Refers to the log-likelihood of the token. The first token of a context will not
	// have a likelihood.
	Likelihood float64 `json:"likelihood"`
}

type LikelihoodOptions struct {
	// The string to compute the log-likelihood of.
	Text string `json:"text"`
}

type LikelihoodResponse struct {
	// The sum of the log-likelihoods of each token in the string.
	Likelihood float64 `json:"likelihood"`

	// An array of token log-likelihood pairs.
	TokenLikelihoods []TokenLikelihood `json:"token_likelihoods"`
}
