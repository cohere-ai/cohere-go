package cohere

type ClassifyOptions struct {
	// An optional string representing the model you'd like to use.
	Model string `json:"model,omitempty"`

	// An optional string representing the ID of a custom playground preset.
	Preset string `json:"preset,omitempty"`

	// An array of strings that you would like to classify.
	Inputs []string `json:"inputs"`

	// An array of ClassifyExamples representing examples and the corresponding label.
	Examples []Example `json:"examples"`

	// Accepts one of NONE, START and END. Determines how inputs over the maximum context length will be handled.
	// Passing START will discard the start of the input and END will discard the end of the input.
	// Defaults to NONE, which will return an error if the input is too long.
	Truncate string `json:"truncate,omitempty"`
}

type Example struct {
	// The text of the example.
	Text string `json:"text"`

	// The label that fits the example's text.
	Label string `json:"label"`
}

type Confidence struct {
	// The label.
	Label string `json:"label"`

	// The associated confidence with the label.
	Confidence float32 `json:"confidence"`
}

type LabelProperties struct {
	Confidence float32 `json:"confidence"`
}

type Classification struct {
	ID string `json:"id"`

	// The top predicted label for the text.
	Prediction string `json:"prediction"`

	// Confidence score for the top predicted label.
	Confidence float32 `json:"confidence"`

	// Confidence score for each label.
	Labels map[string]LabelProperties `json:"labels"`

	// The text that is being classified.
	Input string `json:"input"`
}

type ClassifyResponse struct {
	ID              string           `json:"id"`
	Classifications []Classification `json:"classifications"`

	// Metadata about the API version
	Meta *MetaResponse `json:"meta,omitempty"`
}
