package cohere

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path"
)

type Client struct {
	APIKey  string
	BaseURL string
	Client  http.Client
	Version string
}

const (
	endpointGenerate   = "generate"
	endpointChooseBest = "choose-best"
	endpointEmbed      = "embed"
	endpointLikelihood = "likelihood"

	endpointCheckAPIKey = "check-api-key"
	endpointTokenize    = "tokenize"
)

type CheckAPIKeyResponse struct {
	Valid bool
}

// Public functions

func CreateClient(apiKey string) (*Client, error) {
	client := &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.cohere.ai/",
		Client:  *http.DefaultClient,
		Version: "2021-11-08",
	}

	res, err := client.CheckAPIKey()
	if err != nil {
		return nil, err
	}

	ret := &CheckAPIKeyResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	if !ret.Valid {
		return nil, errors.New("invalid api key")
	}
	return client, nil
}

// Client methods

func (c *Client) post(model string, endpoint string, body interface{}) ([]byte, error) {
	url := c.BaseURL + path.Join(model, endpoint)
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "BEARER "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Request-Source", "go-sdk")
	if len(c.Version) > 0 {
		req.Header.Set("Cohere-Version", c.Version)
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		e := &APIError{}
		if err := json.Unmarshal(buf, e); err != nil {
			return nil, err
		}
		e.StatusCode = res.StatusCode
		return nil, e
	}
	return buf, nil
}

func (c *Client) CheckAPIKey() ([]byte, error) {
	url := c.BaseURL + endpointCheckAPIKey
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "BEARER "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Request-Source", "go-sdk")
	if len(c.Version) > 0 {
		req.Header.Set("Cohere-Version", c.Version)
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		e := &APIError{}
		if err := json.Unmarshal(buf, e); err != nil {
			return nil, err
		}
		e.StatusCode = res.StatusCode
		return nil, e
	}
	return buf, nil
}

// Generates realistic text conditioned on a given input.
// See: https://docs.cohere.ai/generate-reference
// Returns a GenerateResponse object.
func (c *Client) Generate(model string, opts GenerateOptions) (*GenerateResponse, error) {
	if opts.NumGenerations == 0 {
		opts.NumGenerations = 1
	}

	res, err := c.post(model, endpointGenerate, opts)
	if err != nil {
		return nil, err
	}

	ret := &GenerateResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Uses likelihood to perform classification. Given a query text that you'd like to classify between
// a number of options, Choose Best will return a score between the query and each option.
// See: https://docs.cohere.ai/choose-best-reference
// Returns a ChooseBestResponse object.
func (c *Client) ChooseBest(model string, opts ChooseBestOptions) (*ChooseBestResponse, error) {
	res, err := c.post(model, endpointChooseBest, opts)
	if err != nil {
		return nil, err
	}

	ret := &ChooseBestResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Returns text embeddings. An embedding is a list of floating point numbers that captures semantic
// information about the text that it represents.
// See: https://docs.cohere.ai/embed-reference
// Returns an EmbedResponse object.
func (c *Client) Embed(model string, opts EmbedOptions) (*EmbedResponse, error) {
	res, err := c.post(model, endpointEmbed, opts)
	if err != nil {
		return nil, err
	}

	ret := &EmbedResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Provides the model log-likelihood of each token in a string as well as the sum of the log-likelihoods
// of each token in that string.
// See: https://docs.cohere.ai/likelihood-reference
// Returns a LikelihoodResponse object.
func (c *Client) Likelihood(model string, opts LikelihoodOptions) (*LikelihoodResponse, error) {
	res, err := c.post(model, endpointLikelihood, opts)
	if err != nil {
		return nil, err
	}

	ret := &LikelihoodResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) Tokenize(model string, opts TokenizeOptions) (*TokenizeResponse, error) {
	res, err := c.post(model, endpointTokenize, opts)
	if err != nil {
		return nil, err
	}

	ret := &TokenizeResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
