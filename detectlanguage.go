package cohere

type Language struct {
	// Name of the language, eg "French"
	LanguageName string `json:"language_name"`

	// Code of the language, eg "fr"
	LanguageCode string `json:"language_code"`

	// A score between 0 and 1 that represents the confidence of the result.
	Confidence float32 `json:"confidence"`
}

type DetectLanguageOptions struct {
	// Texts to identify languages for
	Texts []string `json:"texts"`
}

type DetectLanguageResponse struct {
	// List of detected languages, one per text
	Results []Language `json:"results"`
}
