// This file was auto-generated by Fern from our API Definition.

package datasets

import (
	context "context"
	v2 "github.com/cohere-ai/cohere-go/v2"
	core "github.com/cohere-ai/cohere-go/v2/core"
	internal "github.com/cohere-ai/cohere-go/v2/internal"
	option "github.com/cohere-ai/cohere-go/v2/option"
	io "io"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.Token == "" {
		options.Token = os.Getenv("CO_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// List datasets that have been created.
func (c *Client) List(
	ctx context.Context,
	request *v2.DatasetsListRequest,
	opts ...option.RequestOption,
) (*v2.DatasetsListResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.cohere.com",
	)
	endpointURL := baseURL + "/v1/datasets"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &v2.TooManyRequestsError{
				APIError: apiError,
			}
		},
		498: func(apiError *core.APIError) error {
			return &v2.InvalidTokenError{
				APIError: apiError,
			}
		},
		499: func(apiError *core.APIError) error {
			return &v2.ClientClosedRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
		501: func(apiError *core.APIError) error {
			return &v2.NotImplementedError{
				APIError: apiError,
			}
		},
		503: func(apiError *core.APIError) error {
			return &v2.ServiceUnavailableError{
				APIError: apiError,
			}
		},
		504: func(apiError *core.APIError) error {
			return &v2.GatewayTimeoutError{
				APIError: apiError,
			}
		},
	}

	var response *v2.DatasetsListResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create a dataset by uploading a file. See ['Dataset Creation'](https://docs.cohere.com/docs/datasets#dataset-creation) for more information.
func (c *Client) Create(
	ctx context.Context,
	data io.Reader,
	evalData io.Reader,
	request *v2.DatasetsCreateRequest,
	opts ...option.RequestOption,
) (*v2.DatasetsCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.cohere.com",
	)
	endpointURL := baseURL + "/v1/datasets"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &v2.TooManyRequestsError{
				APIError: apiError,
			}
		},
		498: func(apiError *core.APIError) error {
			return &v2.InvalidTokenError{
				APIError: apiError,
			}
		},
		499: func(apiError *core.APIError) error {
			return &v2.ClientClosedRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
		501: func(apiError *core.APIError) error {
			return &v2.NotImplementedError{
				APIError: apiError,
			}
		},
		503: func(apiError *core.APIError) error {
			return &v2.ServiceUnavailableError{
				APIError: apiError,
			}
		},
		504: func(apiError *core.APIError) error {
			return &v2.GatewayTimeoutError{
				APIError: apiError,
			}
		},
	}
	writer := internal.NewMultipartWriter()
	if err := writer.WriteFile("data", data); err != nil {
		return nil, err
	}
	if evalData != nil {
		if err := writer.WriteFile("eval_data", evalData); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers.Set("Content-Type", writer.ContentType())

	var response *v2.DatasetsCreateResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         writer.Buffer(),
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// View the dataset storage usage for your Organization. Each Organization can have up to 10GB of storage across all their users.
func (c *Client) GetUsage(
	ctx context.Context,
	opts ...option.RequestOption,
) (*v2.DatasetsGetUsageResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.cohere.com",
	)
	endpointURL := baseURL + "/v1/datasets/usage"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &v2.TooManyRequestsError{
				APIError: apiError,
			}
		},
		498: func(apiError *core.APIError) error {
			return &v2.InvalidTokenError{
				APIError: apiError,
			}
		},
		499: func(apiError *core.APIError) error {
			return &v2.ClientClosedRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
		501: func(apiError *core.APIError) error {
			return &v2.NotImplementedError{
				APIError: apiError,
			}
		},
		503: func(apiError *core.APIError) error {
			return &v2.ServiceUnavailableError{
				APIError: apiError,
			}
		},
		504: func(apiError *core.APIError) error {
			return &v2.GatewayTimeoutError{
				APIError: apiError,
			}
		},
	}

	var response *v2.DatasetsGetUsageResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a dataset by ID. See ['Datasets'](https://docs.cohere.com/docs/datasets) for more information.
func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*v2.DatasetsGetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.cohere.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/datasets/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &v2.TooManyRequestsError{
				APIError: apiError,
			}
		},
		498: func(apiError *core.APIError) error {
			return &v2.InvalidTokenError{
				APIError: apiError,
			}
		},
		499: func(apiError *core.APIError) error {
			return &v2.ClientClosedRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
		501: func(apiError *core.APIError) error {
			return &v2.NotImplementedError{
				APIError: apiError,
			}
		},
		503: func(apiError *core.APIError) error {
			return &v2.ServiceUnavailableError{
				APIError: apiError,
			}
		},
		504: func(apiError *core.APIError) error {
			return &v2.GatewayTimeoutError{
				APIError: apiError,
			}
		},
	}

	var response *v2.DatasetsGetResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete a dataset by ID. Datasets are automatically deleted after 30 days, but they can also be deleted manually.
func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (map[string]interface{}, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.cohere.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/v1/datasets/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &v2.TooManyRequestsError{
				APIError: apiError,
			}
		},
		498: func(apiError *core.APIError) error {
			return &v2.InvalidTokenError{
				APIError: apiError,
			}
		},
		499: func(apiError *core.APIError) error {
			return &v2.ClientClosedRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
		501: func(apiError *core.APIError) error {
			return &v2.NotImplementedError{
				APIError: apiError,
			}
		},
		503: func(apiError *core.APIError) error {
			return &v2.ServiceUnavailableError{
				APIError: apiError,
			}
		},
		504: func(apiError *core.APIError) error {
			return &v2.GatewayTimeoutError{
				APIError: apiError,
			}
		},
	}

	var response map[string]interface{}
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
