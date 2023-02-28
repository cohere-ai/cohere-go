package cohere

type SummarizeOptions struct {
	// Text to summarize
	Text string `json:"text"`
	// 'One of `paragraph` or `bullets`, defaults to `paragraph`.
	// Indicates the style in which the summary will be delivered - in a free form
	// paragraph or in bullet points.'
	Format string `json:"format,omitempty"`
	// One of `short`, `medium` or `long`, defaults to `medium`. Indicates the approximate length of the summary.'
	Length string `json:"length,omitempty"`
	// One of `low`, `medium` or `high`, defaults to `low`. Controls how close to the original text the summary is.
	// `high` extractiveness summaries will lean towards reusing sentences verbatim, while `low` extractiveness
	// summaries will tend to paraphrase more.'
	Extractiveness string `json:"string,omitempty"`
	// Ranges from 0 to 5. Controls the randomness of the output. Lower values tend to generate more “predictable” output,
	// while higher values tend to generate more “creative” output. The sweet spot is typically between 0 and 1.
	Temperature *float64 `json:"temperature,omitempty"`
	// A free-form instruction for modifying how the summaries get generated. Should complete the sentence "Generate a summary _".
	// Eg. "focusing on the next steps" or "written by Yoda"
	AdditionalCommand string `json:"additional_command,omitempty"`
	// Denotes the summarization model to be used. Defaults to the best performing model
	Model string `json:"model,omitempty"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
	ID      string `json:"id"`

	// Metadata about the API version
	Meta *MetaResponse `json:"meta,omitempty"`
}
