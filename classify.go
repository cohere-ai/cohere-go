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

type LabelPrediction struct {
	// The associated confidence with the label.
	Confidence float32 `json:"confidence"`
}
type Classification struct {
	// The text that is being classified.
	Input string `json:"input"`

	// The predicted labels and confidences for the text.
	Prediction map[string]float32 `json:"prediction"`

	// Confidence score for the predicted label.
	Confidence float32 `json:"confidence"`

	// Confidence scores for each label.
	Labels map[string]float32 `json:"labels"`

	// The confidence score for each label.
	Labels map[string]LabelPrediction `json:"labels"`

	// The predicted label for the text.
	PredictionLabel string

	// The confidence for the predicted label.
	PredictionConfidence float32
}

type ClassifyResponse struct {
	Classifications []Classification `json:"classifications"`
}
