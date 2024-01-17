// This file was auto-generated by Fern from our API Definition.

package api

type CreateConnectorRequest struct {
	// A human-readable name for the connector.
	Name string `json:"name"`
	// A description of the connector.
	Description *string `json:"description,omitempty"`
	// The URL of the connector that will be used to search for documents.
	Url string `json:"url"`
	// A list of fields to exclude from the prompt (fields remain in the document).
	Excludes []string `json:"excludes,omitempty"`
	// The OAuth 2.0 configuration for the connector. Cannot be specified if service_auth is specified.
	Oauth *CreateConnectorOAuth `json:"oauth,omitempty"`
	// Whether the connector is active or not.
	Active *bool `json:"active,omitempty"`
	// Whether a chat request should continue or not if the request to this connector fails.
	ContinueOnFailure *bool `json:"continue_on_failure,omitempty"`
	// The service to service authentication configuration for the connector. Cannot be specified if oauth is specified.
	ServiceAuth *CreateConnectorServiceAuth `json:"service_auth,omitempty"`
}

type ConnectorsListRequest struct {
	// Maximum number of connectors to return [0, 100].
	Limit *float64 `json:"-"`
	// Number of connectors to skip before returning results [0, inf].
	Offset *float64 `json:"-"`
}

type ConnectorsOAuthAuthorizeRequest struct {
	// The URL to redirect to after the connector has been authorized.
	AfterTokenRedirect *string `json:"-"`
}

type UpdateConnectorRequest struct {
	// A human-readable name for the connector.
	Name *string `json:"name,omitempty"`
	// The URL of the connector that will be used to search for documents.
	Url *string `json:"url,omitempty"`
	// A list of fields to exclude from the prompt (fields remain in the document).
	Excludes []string `json:"excludes,omitempty"`
	// The OAuth 2.0 configuration for the connector. Cannot be specified if service_auth is specified.
	Oauth             *CreateConnectorOAuth `json:"oauth,omitempty"`
	Active            *bool                 `json:"active,omitempty"`
	ContinueOnFailure *bool                 `json:"continue_on_failure,omitempty"`
	// The service to service authentication configuration for the connector. Cannot be specified if oauth is specified.
	ServiceAuth *CreateConnectorServiceAuth `json:"service_auth,omitempty"`
}
