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
	// An optional string representing the model you'd like to use.
	Model string `json:"model,omitempty"`

	// Represents the prompt or text to be completed.
	Prompt string `json:"prompt,omitempty"`

	// optional - Denotes the number of tokens to predict per generation.
	MaxTokens *uint `json:"max_tokens,omitempty"`

	// optional - The ID of a custom playground preset.
	Preset string `json:"preset,omitempty"`

	// optional - A non-negative float that tunes the degree of randomness in generation.
	Temperature *float64 `json:"temperature,omitempty"`

	// optional - Denotes the maximum number of generations that will be returned. Defaults to 1,
	// max value of 5.
	NumGenerations *int `json:"num_generations,omitempty"`

	// optional - If set to a positive integer, it ensures only the top k most likely tokens are
	// considered for generation at each step.
	K *int `json:"k,omitempty"`

	// optional - If set to a probability 0.0 < p < 1.0, it ensures that only the most likely tokens,
	// with total probability mass of p, are considered for generation at each step. If both k and
	// p are enabled, p acts after k. Max value of 1.0.
	P *float64 `json:"p,omitempty"`

	// optional - Can be used to reduce repetitiveness of generated tokens. The higher the value,
	// the stronger a penalty is applied to previously present tokens, proportional to how many
	// times they have already appeared in the prompt or prior generation. Max value of 1.0.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// optional - Can be used to reduce repetitiveness of generated tokens. Similar to frequency_penalty,
	// except that this penalty is applied equally to all tokens that have already appeared, regardless
	// of their exact frequencies. Max value of 1.0.
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`

	// optional - The generated text will be cut at the beginning of the earliest occurence of an end sequence.
	// The sequence will be excluded from the text.
	EndSequences []string `json:"end_sequences,omitempty"`

	// optional - The generated text will be cut at the end of the earliest occurence of a stop sequence.
	// The sequence will be included the text.
	StopSequences []string `json:"stop_sequences,omitempty"`

	// optional - One of GENERATION|ALL|NONE to specify how and if the token likelihoods are returned with
	// the response. If GENERATION is selected, the token likelihoods will only be provided for generated
	// text. If ALL is selected, the token likelihoods will be provided both for the prompt and the generated
	// text.
	ReturnLikelihoods string `json:"return_likelihoods,omitempty"`

	// optional - Used to prevent the model from generating unwanted tokens or to incentivize it to include desired tokens
	// A map of tokens to biases where bias is a float between -10 and +10
	// Negative values will disincentivize that token from appearing while positivse values will incentivize them
	// Tokens can be obtained from text using the tokenizer
	// Note: logit bias may not be supported for all finetune models
	LogitBias map[int]float32 `json:"logit_bias,omitempty"`
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
