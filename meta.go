package cohere

type MetaResponse struct {
	APIVersion *APIVersionMeta `json:"api_version,omitempty"`
	Warnings   []string        `json:"warnings,omitempty"`
}

// Metadata about the API version being used
type APIVersionMeta struct {
	Version        string `json:"version"`
	IsDeprecated   bool   `json:"is_deprecated,omitempty"`
	IsExperimental bool   `json:"is_experimental,omitempty"`
}
