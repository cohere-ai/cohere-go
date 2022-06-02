package cohere

// return likelihoods
const (
	ReturnGeneration = "GENERATION"
	ReturnAll        = "ALL"
	ReturnNone       = "NONE"
)

type TokenLikelihood struct {
	// The token.
	Token string `json:"token"`

	// Refers to the log-likelihood of the token. The first token of a context will not
	// have a likelihood.
	Likelihood float64 `json:"likelihood"`
}

type GenerateOptions struct {
	// Represents the prompt or text to be completed.
	Prompt string `json:"prompt"`

	// Denotes the number of tokens to predict per generation.
	MaxTokens uint `json:"max_tokens"`

	// A non-negative float that tunes the degree of randomness in generation.
	Temperature float64 `json:"temperature"`

	// optional - Denotes the maximum number of generations that will be returned. Defaults to 1,
	// max value of 5.
	NumGenerations int `json:"num_generations"`

	// optional - If set to a positive integer, it ensures only the top k most likely tokens are
	// considered for generation at each step.
	K int `json:"k"`

	// optional - If set to a probability 0.0 < p < 1.0, it ensures that only the most likely tokens,
	// with total probability mass of p, are considered for generation at each step. If both k and
	// p are enabled, p acts after k. Max value of 1.0.
	P float64 `json:"p"`

	// optional - Can be used to reduce repetitiveness of generated tokens. The higher the value,
	// the stronger a penalty is applied to previously present tokens, proportional to how many
	// times they have already appeared in the prompt or prior generation. Max value of 1.0.
	FrequencyPenalty float64 `json:"frequency_penalty"`

	// optional - Can be used to reduce repetitiveness of generated tokens. Similar to frequency_penalty,
	// except that this penalty is applied equally to all tokens that have already appeared, regardless
	// of their exact frequencies. Max value of 1.0.
	PresencePenalty float64 `json:"presence_penalty"`

	// optional - A stop sequence will cut off your generation at the end of the sequence. Providing multiple
	// stop sequences in the array will cut the generation at the first stop sequence in the generation,
	// if applicable.
	StopSequences []string `json:"stop_sequences,omitempty"`

	// optional - One of GENERATION|ALL|NONE to specify how and if the token likelihoods are returned with
	// the response. If GENERATION is selected, the token likelihoods will only be provided for generated
	// text. If ALL is selected, the token likelihoods will be provided both for the prompt and the generated
	// text.
	ReturnLikelihoods string `json:"return_likelihoods,omitempty"`

	// optional - Language code (eg: "fr" for French) to specify the language of the generated text.
	// Support for languages varies between models. By default, the language is set to English.
	Language string `json:"language,omitempty"`
}

type Generation struct {
	// Contains the generated text.
	Text string `json:"text"`

	// The sum of the log-likehoods of each token in the string.
	Likelihood *float64 `json:"likelihood,omitempty"`

	// Only returned if `return_likelihoods` is not set to NONE.
	// The likelihood.
	TokenLikelihoods []TokenLikelihood `json:"token_likelihoods,omitempty"`
}

type GenerateResponse struct {
	// Contains the generations.
	Generations []Generation `json:"generations"`
}
