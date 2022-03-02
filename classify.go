package cohere

type ClassifyOptions struct {
	// An optional string representing what you'd like the model to do.
	Task string `json:"task,omitempty"`

	// An array of strings that you would like to classify.
	Texts []string `json:"texts"`

	// An array of ClassifyExamples representing examples and the corresponding label.
	Examples []Example `json:"examples"`

	// An optional string to append onto every example and text prior to the label.
	Prompt string `json:"prompt,omitempty"`
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
type Classification struct {
	// The text that is being classified.
	Text string `json:"text"`

	// The predicted label for the text.
	Prediction string `json:"prediction"`

	// The confidence score for each label.
	Confidences []Confidence `json:"confidences"`
}

type ClassifyResponse struct {
	Classifications []Classification `json:"results"`
}
