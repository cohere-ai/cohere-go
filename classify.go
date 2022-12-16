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

type LabelConfidence struct {
	Confidence float32 `json:"confidence"`
}

type Classification struct {

	// The top predicted label for the text.
	Prediction string `json:"prediction"`

	// Confidence score for the top predicted label.
	Confidence float32 `json:"confidence"`

	// Confidence score for each label.
	Labels map[string]LabelConfidence `json:"labels"`

	// The text that is being classified.
	Input string `json:"input"`
}

type ClassifyResponse struct {
	Classifications []Classification `json:"classifications"`
}
