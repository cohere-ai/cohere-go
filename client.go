package cohere

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cohere-ai/tokenizer"
)

type Client struct {
	APIKey  string
	BaseURL string
	Client  http.Client
	version string
}

const (
	endpointGenerate       = "generate"
	endpointEmbed          = "embed"
	endpointClassify       = "classify"
	endpointDetectLanguage = "detect-language"
	endpointSummarize      = "summarize"
	endpointRerank         = "rerank"

	// Truncate modes for co.embed, co.generate and co.classify
	NONE  = "NONE"
	START = "START"
	END   = "END"

	endpointCheckAPIKey = "check-api-key"
)

type CheckAPIKeyResponse struct {
	Valid bool
}

// Public functions

func CreateClient(apiKey string) (*Client, error) {
	client := &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.cohere.ai",
		Client:  *http.DefaultClient,
		version: "v1",
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

func (c *Client) post(endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, c.version, endpoint)
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", c.APIKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Request-Source", "go-sdk")

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

	for _, warning := range res.Header.Values("X-API-Warning") {
		fmt.Fprintf(os.Stderr, "\033[93mWarning: %s\n\033[0m", warning)
	}
	return buf, nil
}

func (c *Client) CheckAPIKey() ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpointCheckAPIKey)
	req, err := http.NewRequest("POST", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", c.APIKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Request-Source", "go-sdk")

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
func (c *Client) Generate(opts GenerateOptions) (*GenerateResponse, error) {
	res, err := c.post(endpointGenerate, opts)
	if err != nil {
		return nil, err
	}

	ret := &GenerateResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Stream streams realistic text conditioned on a given input.
// Callers must examine the GenerationResult.Err field to
// determine if an error occurred. There could be multiple
// errors in the stream: one per requested generation,
// see GenerateOptions.NumGenerations.
//
// Note: this func will close channel once response is exhausted.
func (c *Client) Stream(opts GenerateOptions) <-chan *GenerationResult {
	ch := make(chan *GenerationResult)

	go func() {
		defer close(ch)

		url := fmt.Sprintf("%s/%s", c.BaseURL, endpointGenerate)
		opts.Stream = true
		buf, err := json.Marshal(opts)
		if err != nil {
			ch <- &GenerationResult{Err: err}
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
		if err != nil {
			ch <- &GenerationResult{Err: err}
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", c.APIKey))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Request-Source", "go-sdk")
		res, err := c.Client.Do(req)
		if err != nil {
			ch <- &GenerationResult{Err: err}
			return
		}
		defer res.Body.Close()

		if res.StatusCode < 200 || res.StatusCode >= 300 {
			ch <- &GenerationResult{Err: fmt.Errorf("HTTP status: %v", res.StatusCode)}
			return
		}

		dec := json.NewDecoder(res.Body)
		for {
			msg := &GeneratedToken{}
			if err := dec.Decode(msg); err != nil {
				if err == io.EOF {
					break
				}
				ch <- &GenerationResult{
					Err: err,
				}
				break
			}
			ch <- &GenerationResult{
				Token: msg,
			}
		}
	}()
	return ch
}

// Classifies text as one of the given labels. Returns a confidence score for each label.
// See: https://docs.cohere.ai/classify-reference
// Returns a ClassifyResponse object.
func (c *Client) Classify(opts ClassifyOptions) (*ClassifyResponse, error) {
	res, err := c.post(endpointClassify, opts)
	if err != nil {
		return nil, err
	}

	ret := &ClassifyResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Returns text embeddings. An embedding is a list of floating point numbers that captures semantic
// information about the text that it represents.
// See: https://docs.cohere.ai/embed-reference
// Returns an EmbedResponse object.
func (c *Client) Embed(opts EmbedOptions) (*EmbedResponse, error) {
	res, err := c.post(endpointEmbed, opts)
	if err != nil {
		return nil, err
	}

	ret := &EmbedResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Tokenizes a string.
// Returns a TokenizeResponse object.
func (c *Client) Tokenize(opts TokenizeOptions) (*TokenizeResponse, error) {
	return Tokenize(opts)
}

func Tokenize(opts TokenizeOptions) (*TokenizeResponse, error) {
	encoder, err := tokenizer.NewFromPrebuilt("coheretext-50k")
	if err != nil {
		return nil, err
	}
	tokens, tokenStrings := encoder.Encode(opts.Text)
	ret := &TokenizeResponse{
		Tokens:       tokens,
		TokenStrings: tokenStrings,
	}
	return ret, nil
}

// Returns a string that corresponds to the provided tokens.
// Returns a DetokenizeResponse object.
func (c *Client) Detokenize(opts DetokenizeOptions) (*DetokenizeResponse, error) {
	return Detokenize(opts)
}

func Detokenize(opts DetokenizeOptions) (*DetokenizeResponse, error) {
	encoder, err := tokenizer.NewFromPrebuilt("coheretext-50k")
	if err != nil {
		return nil, err
	}
	text := encoder.Decode(opts.Tokens)
	ret := &DetokenizeResponse{
		Text: text,
	}
	return ret, nil
}

// For each of the provided texts, returns the expected language of that text.
// See: https://docs.cohere.ai/detect-language-reference
// Returns a DetectLanguageResponse object.
func (c *Client) DetectLanguage(opts DetectLanguageOptions) (*DetectLanguageResponse, error) {
	res, err := c.post(endpointDetectLanguage, opts)
	if err != nil {
		return nil, err
	}

	ret := &DetectLanguageResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) Summarize(opts SummarizeOptions) (*SummarizeResponse, error) {
	res, err := c.post(endpointSummarize, opts)
	if err != nil {
		return nil, err
	}

	ret := &SummarizeResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// For a query and a list of texts, returns a ordered array of relevance scores for each text.
// See: https://docs.cohere.ai/reference/rerank-1
// Returns a RerankResponse object.
func (c *Client) Rerank(options RerankOptions) (*RerankResponse, error) {
	res, err := c.post(endpointRerank, options)
	if err != nil {
		return nil, err
	}

	ret := &RerankResponse{}
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
