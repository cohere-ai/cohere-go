package cohere

type SummarizeOptions struct {
	// Texts to identify languages for
	Text string `json:"text"`
	// One of "paragraph" or "bullets"
	Format string `json:"format,omitempty"`
	// One of "short", "medium" or "long"
	Length string `json:"length,omitempty"`
	// Number between 0 and 1 which indicates how much randomness is applied between the generations
	Temperature *float64 `json:"temperature,omitempty"`
	// Modifies the underlying prompt. Completes the sentence "Generate a summary _"
	AdditionalCommand string `json:"additional_command,omitempty"`
	// The summarization model to use
	Model string `json:"model,omitempty"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
	ID      string `json:"id"`
}
