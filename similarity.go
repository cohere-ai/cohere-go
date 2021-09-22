package cohere

type SimilarityOptions struct {
	// Used for comparison against the strings in `targets`.
	Anchor string `json:"anchor"`

	// To be compared to anchor.
	Targets []string `json:"targets"`
}

type SimilarityResponse struct {
	// An array of floats, where each entry represents the similarity of each target to the
	// anchor respectively. A higher value means that a target is more similar to the anchor.
	Similarities []float64 `json:"similarities"`
}
