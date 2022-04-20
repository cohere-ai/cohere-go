package cohere

type ClassifyOptions struct {
	// An optional string representing what you'd like the model to do.
	TaskDescription string `json:"taskDescription,omitempty"`

	// An array of strings that you would like to classify.
	Inputs []string `json:"inputs"`

	// An array of ClassifyExamples representing examples and the corresponding label.
	// Required when using baseline models, but can be ommitted when using classification finetuned models.
	// See https://docs.cohere.ai/classify-reference for more details.
	Examples []Example `json:"examples,omitempty"`

	// An optional string to append onto every example and text prior to the label.
	OutputIndicator string `json:"outputIndicator,omitempty"`
}

type Example struct {
	// The text of the example.
	Text string `json:"text"`

	// The class that fits the example's text.
	Class string `json:"class"`
}

type Confidence struct {
	// The class.
	Class string `json:"class"`

	// The associated confidence value with the label.
	Value float32 `json:"value"`
}
type Classification struct {
	// The text that is being classified.
	Input string `json:"input"`

	// The predicted label for the text.
	Prediction string `json:"prediction"`

	// The confidence score for each label.
	Confidences []Confidence `json:"confidences"`
}

type ClassifyResponse struct {
	Classifications []Classification `json:"classifications"`
}
