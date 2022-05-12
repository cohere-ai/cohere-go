package cohere

type ExtractOptions struct {
	// An array of strings that you would like to run extraction on.
	Texts []string `json:"texts"`

	// An array of ExtractExamples representing examples and the corresponding entities.
	Examples []ExtractExample `json:"examples"`
}

type ExtractExample struct {
	// The text of the example.
	Text string `json:"text"`

	// The label that fits the example's text.
	Entities []ExtractEntity `json:"entities"`
}

type ExtractEntity struct {
	// The type of the extracted entity, eg: "Food"
	Type string `json:"type"`

	// The value of the extracted entity, eg: "Pizza"
	Value string `json:"value"`
}

type Extraction struct {
	// Id of the performed extraction.
	ID string `json:"id"`

	// The input text.
	Text string `json:"text"`

	// An array of extracted entities.
	Entities []ExtractEntity `json:"entities"`
}

type ExtractResponse struct {
	Extractions []Extraction `json:"results"`
}
