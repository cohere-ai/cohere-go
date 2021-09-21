package cohere

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Model string

type Client struct {
	apiKey string
	client http.Client
}

// Models
const (
	Shrimp Model = "baseline-shrimp"
	Otter  Model = "baseline-otter"
	Seal   Model = "baseline-seal"
	Shark  Model = "baseline-shark"
	Orca   Model = "baseline-orca"
)

const baseUrl = "https://api.cohere.ai/"
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
		apiKey: apiKey,
		client: *http.DefaultClient,
	}
}

// Client methods

func (c *Client) post(model Model, endpoint string, body interface{}) ([]byte, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	url := baseUrl + string(model) + endpoint
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "BEARER "+c.apiKey)
	res, err := c.client.Do(req)
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
func (c *Client) Generate(model Model, prompt string, maxTokens uint, temperature float64) (string, error) {
	body := GenerateOptions{
		Prompt:            prompt,
		MaxTokens:         maxTokens,
		Temperature:       temperature,
		K:                 0,
		P:                 0.75,
		FrequencyPenalty:  0.0,
		PresencePenalty:   0.0,
		StopSequences:     []string{},
		ReturnLikelihoods: NONE,
	}

	res, err := c.post(model, endpointGenerate, body)
	if err != nil {
		return "", err
	}

	ret := &GenerateResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return "", err
	}
	return ret.Text, nil
}

// Generates realistic text conditioned on a given input with advanced configuration.
// See: https://docs.cohere.ai/generate-reference
func (c *Client) GenerateAdvanced(model Model, opts GenerateOptions) (*GenerateResponse, error) {
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
func (c *Client) Similarity(model Model, anchor string, targets []string) ([]float64, error) {
	body := SimilarityRequest{
		Anchor:  anchor,
		Targets: targets,
	}

	res, err := c.post(model, endpointSimilarity, body)
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
func (c *Client) ChooseBest(model Model, query string, options []string, mode string) (*ChooseBestResponse, error) {
	body := ChooseBestRequest{
		Query:   query,
		Options: options,
		Mode:    mode,
	}

	res, err := c.post(model, endpointChooseBest, body)
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
func (c *Client) Embed(model Model, texts []string) ([][]float64, error) {
	body := EmbedRequest{
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
func (c *Client) Likelihood(model Model, text string) (*LikelihoodResponse, error) {
	body := LikelihoodRequest{
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
