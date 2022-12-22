package cohere

type LanguageDetectResult struct {
	// Name of the language, eg "French"
	LanguageName string `json:"language_name"`

	// Code of the language, eg "fr"
	LanguageCode string `json:"language_code"`
}

type DetectLanguageOptions struct {
	// Texts to identify languages for
	Texts []string `json:"texts"`
}

type DetectLanguageResponse struct {
	// List of detected languages, one per text
	Results []LanguageDetectResult `json:"results"`
}
