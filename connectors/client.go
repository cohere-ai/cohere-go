// This file was auto-generated by Fern from our API Definition.

package connectors

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	v2 "github.com/cohere-ai/cohere-go/v2"
	core "github.com/cohere-ai/cohere-go/v2/core"
	option "github.com/cohere-ai/cohere-go/v2/option"
	io "io"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.Token == "" {
		options.Token = os.Getenv("CO_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Returns a list of connectors ordered by descending creation date (newer first). See ['Managing your Connector'](https://docs.cohere.com/docs/managing-your-connector) for more information.
func (c *Client) List(
	ctx context.Context,
	request *v2.ConnectorsListRequest,
	opts ...option.RequestOption,
) (*v2.ListConnectorsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/connectors"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v2.ListConnectorsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Creates a new connector. The connector is tested during registration and will cancel registration when the test is unsuccessful. See ['Creating and Deploying a Connector'](https://docs.cohere.com/v1/docs/creating-and-deploying-a-connector) for more information.
func (c *Client) Create(
	ctx context.Context,
	request *v2.CreateConnectorRequest,
	opts ...option.RequestOption,
) (*v2.CreateConnectorResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/connectors"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v2.CreateConnectorResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a connector by ID. See ['Connectors'](https://docs.cohere.com/docs/connectors) for more information.
func (c *Client) Get(
	ctx context.Context,
	// The ID of the connector to retrieve.
	id string,
	opts ...option.RequestOption,
) (*v2.GetConnectorResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/connectors/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v2.GetConnectorResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete a connector by ID. See ['Connectors'](https://docs.cohere.com/docs/connectors) for more information.
func (c *Client) Delete(
	ctx context.Context,
	// The ID of the connector to delete.
	id string,
	opts ...option.RequestOption,
) (v2.DeleteConnectorResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/connectors/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response v2.DeleteConnectorResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update a connector by ID. Omitted fields will not be updated. See ['Managing your Connector'](https://docs.cohere.com/docs/managing-your-connector) for more information.
func (c *Client) Update(
	ctx context.Context,
	// The ID of the connector to update.
	id string,
	request *v2.UpdateConnectorRequest,
	opts ...option.RequestOption,
) (*v2.UpdateConnectorResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/connectors/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v2.UpdateConnectorResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Authorize the connector with the given ID for the connector oauth app. See ['Connector Authentication'](https://docs.cohere.com/docs/connector-authentication) for more information.
func (c *Client) OAuthAuthorize(
	ctx context.Context,
	// The ID of the connector to authorize.
	id string,
	request *v2.ConnectorsOAuthAuthorizeRequest,
	opts ...option.RequestOption,
) (*v2.OAuthAuthorizeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.cohere.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/connectors/%v/oauth/authorize", id)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v2.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(v2.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(v2.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(v2.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(v2.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(v2.TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 499:
			value := new(v2.ClientClosedRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(v2.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 501:
			value := new(v2.NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 503:
			value := new(v2.ServiceUnavailableError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 504:
			value := new(v2.GatewayTimeoutError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v2.OAuthAuthorizeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
