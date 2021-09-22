package cohere

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	ApiKey  string
	BaseURL string
	Client  http.Client
}

const (
	endpointGenerate   = "/generate"
	endpointSimilarity = "/similarity"
	endpointChooseBest = "/choose-best"
	endpointEmbed      = "/embed"
	endpointLikelihood = "/likelihood"
)

// Public functions

func CreateClient(apiKey string) *Client {
	return &Client{
		ApiKey:  apiKey,
		BaseURL: "https://api.cohere.ai/",
		Client:  *http.DefaultClient,
	}
}

// Client methods

func (c *Client) post(model string, endpoint string, body interface{}) ([]byte, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	url := c.BaseURL + string(model) + endpoint
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "BEARER "+c.ApiKey)
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
		e := &ApiError{}
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
// Returns an object containing the generated text.
func (c *Client) Generate(model string, opts GenerateOptions) (*GenerateResponse, error) {
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

// Uses embeddings to measure the semantic similarity between a text `anchor` and its `targets`.
// See: https://docs.cohere.ai/similarity-reference
// Returns an array of floats representing the similarity of each target to the anchor.
func (c *Client) Similarity(model string, opts SimilarityOptions) ([]float64, error) {
	res, err := c.post(model, endpointSimilarity, opts)
	if err != nil {
		return nil, err
	}

	ret := &SimilarityResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret.Similarities, nil
}

// Uses likelihood to perform classification. Given a query text that you'd like to classify between
// a number of options, Choose Best will return a score between the query and each option.
// See: https://docs.cohere.ai/choose-best-reference
// Returns an object containing the tokens and score of each token.
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
// Returns an array of embeddings, where each embedding is an array of floats.
func (c *Client) Embed(model string, texts []string) ([][]float64, error) {
	body := EmbedOptions{
		Texts: texts,
	}

	res, err := c.post(model, endpointEmbed, body)
	if err != nil {
		return nil, err
	}

	ret := &EmbedResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret.Embeddings, nil
}

// Provides the model log-likelihood of each token in a string as well as the sum of the log-likelihoods
// of each token in that string.
// See: https://docs.cohere.ai/likelihood-reference
// Returns an object containing the sum and per-token log-likelihoods.
func (c *Client) Likelihood(model string, text string) (*LikelihoodResponse, error) {
	body := LikelihoodOptions{
		Text: text,
	}

	res, err := c.post(model, endpointLikelihood, body)
	if err != nil {
		return nil, err
	}

	ret := &LikelihoodResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
