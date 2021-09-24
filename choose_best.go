package cohere

// mode
const (
	// The output will be the log-likelihood of the `query` conditioned on the `option`: log p(query|option)
	PrependOption = "PREPEND_OPTION"
	// The output will be the log-likelihood of the `option` conditioned on the `query`: log p(option|query)
	AppendOption = "APPEND_OPTION"
)

type ChooseBestOptions struct {
	// Used to query the options.
	Query string `json:"query"`

	// Each string concatenates to the query.
	Options []string `json:"options"`

	// One of PREPEND_OPTION|APPEND_OPTION to specify where the option string will be placed and
	// how to compute the log-likelihood.
	Mode string `json:"mode"`
}

type ChooseBestResponse struct {
	// An array of floats corresponding to a score for each of the options, where a higher score
	// represents a more likely query-option pair.
	Scores []float64 `json:"scores"`

	// An array of tokens corresponding to the tokens for each of the options.
	Tokens [][]string `json:"tokens"`

	// An array of log likelihoods corresponding to the tokens of each of the options.
	TokenLogLikelihoods [][]float64 `json:"token_log_likelihoods"`
}
