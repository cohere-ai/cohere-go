// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/cohere-ai/cohere-go/v2/core"
)

// This error is returned when the request is not well formed. This could be because:
//   - JSON is invalid
//   - The request is missing required fields
//   - The request contains an invalid combination of fields
type BadRequestError struct {
	*core.APIError
	Body interface{}
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// This error is returned when a request is cancelled by the user.
type ClientClosedRequestError struct {
	*core.APIError
	Body interface{}
}

func (c *ClientClosedRequestError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 499
	c.Body = body
	return nil
}

func (c *ClientClosedRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ClientClosedRequestError) Unwrap() error {
	return c.APIError
}

// This error indicates that the operation attempted to be performed is not allowed. This could be because:
//   - The api token is invalid
//   - The user does not have the necessary permissions
type ForbiddenError struct {
	*core.APIError
	Body interface{}
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

// This error is returned when a request to the server times out. This could be due to:
//   - An internal services taking too long to respond
type GatewayTimeoutError struct {
	*core.APIError
	Body interface{}
}

func (g *GatewayTimeoutError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	g.StatusCode = 504
	g.Body = body
	return nil
}

func (g *GatewayTimeoutError) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Body)
}

func (g *GatewayTimeoutError) Unwrap() error {
	return g.APIError
}

// This error is returned when an uncategorised internal server error occurs.
type InternalServerError struct {
	*core.APIError
	Body interface{}
}

func (i *InternalServerError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.StatusCode = 500
	i.Body = body
	return nil
}

func (i *InternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}

func (i *InternalServerError) Unwrap() error {
	return i.APIError
}

// This error is returned when a request or response contains a deny-listed token.
type InvalidTokenError struct {
	*core.APIError
	Body interface{}
}

func (i *InvalidTokenError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.StatusCode = 498
	i.Body = body
	return nil
}

func (i *InvalidTokenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}

func (i *InvalidTokenError) Unwrap() error {
	return i.APIError
}

// This error is returned when a resource is not found. This could be because:
//   - The endpoint does not exist
//   - The resource does not exist eg model id, dataset id
type NotFoundError struct {
	*core.APIError
	Body interface{}
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// This error is returned when the requested feature is not implemented.
type NotImplementedError struct {
	*core.APIError
	Body interface{}
}

func (n *NotImplementedError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 501
	n.Body = body
	return nil
}

func (n *NotImplementedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotImplementedError) Unwrap() error {
	return n.APIError
}

// This error is returned when the service is unavailable. This could be due to:
//   - Too many users trying to access the service at the same time
type ServiceUnavailableError struct {
	*core.APIError
	Body interface{}
}

func (s *ServiceUnavailableError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.StatusCode = 503
	s.Body = body
	return nil
}

func (s *ServiceUnavailableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

func (s *ServiceUnavailableError) Unwrap() error {
	return s.APIError
}

// Too many requests
type TooManyRequestsError struct {
	*core.APIError
	Body interface{}
}

func (t *TooManyRequestsError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	t.StatusCode = 429
	t.Body = body
	return nil
}

func (t *TooManyRequestsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Body)
}

func (t *TooManyRequestsError) Unwrap() error {
	return t.APIError
}

// This error indicates that the operation attempted to be performed is not allowed. This could be because:
//   - The api token is invalid
//   - The user does not have the necessary permissions
type UnauthorizedError struct {
	*core.APIError
	Body interface{}
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

// This error is returned when the request is not well formed. This could be because:
//   - JSON is invalid
//   - The request is missing required fields
//   - The request contains an invalid combination of fields
type UnprocessableEntityError struct {
	*core.APIError
	Body interface{}
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}
