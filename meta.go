package cohere

type MetaResponse struct {
	APIVersion *APIVersionMeta `json:"api_version"`
}

// Metadata about the API version being used
type APIVersionMeta struct {
	Version        string `json:"version"`
	IsDeprecated   bool   `json:"is_deprecated,omitempty"`
	IsExperimental bool   `json:"is_experimental,omitempty"`
}
