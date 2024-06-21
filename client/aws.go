package client

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	errors "errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	option "github.com/cohere-ai/cohere-go/v2/option"
)

func newAwsRequestOptions(opts ...AwsRequestOption) *AwsRequestOptions {
	options := &AwsRequestOptions{}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

type awsClient struct {
	AwsRequestOptions
	Service string
}

// RequestOption adapts the behavior of the client or an individual request.
type AwsRequestOption interface {
	applyRequestOptions(*AwsRequestOptions)
}

type AwsRequestOptions struct {
	environment     string
	awsRegion       string
	awsAccessKey    string
	awsSecretKey    string
	awsSessionToken string
}

// AwsEnvironment implements the RequestOption interface.
type AwsEnvironment struct {
	environment string
}

func (h *AwsEnvironment) applyRequestOptions(opts *AwsRequestOptions) {
	opts.environment = h.environment
}

type AwsRegion struct {
	awsRegion string
}

func (h *AwsRegion) applyRequestOptions(opts *AwsRequestOptions) {
	opts.awsRegion = h.awsRegion
}

func WithAwsRegion(region string) *AwsRegion {
	return &AwsRegion{awsRegion: region}
}

type AwsAccessKey struct {
	awsAccessKey string
}

func (h *AwsAccessKey) applyRequestOptions(opts *AwsRequestOptions) {
	opts.awsAccessKey = h.awsAccessKey
}

func WithAwsAccessKey(accessKey string) *AwsAccessKey {
	return &AwsAccessKey{awsAccessKey: accessKey}
}

type AwsSecretKey struct {
	awsSecretKey string
}

func (h *AwsSecretKey) applyRequestOptions(opts *AwsRequestOptions) {
	opts.awsSecretKey = h.awsSecretKey
}

func WithAwsSecretKey(secretKey string) *AwsSecretKey {
	return &AwsSecretKey{awsSecretKey: secretKey}
}

type AwsSessionToken struct {
	awsSessionToken string
}

func (h *AwsSessionToken) applyRequestOptions(opts *AwsRequestOptions) {
	opts.awsSessionToken = h.awsSessionToken
}

func WithAwsSessionToken(sessionToken string) *AwsSessionToken {
	return &AwsSessionToken{awsSessionToken: sessionToken}
}

func NewAwsClient(baseOpts []option.RequestOption, awsOpts []AwsRequestOption, service string) *Client {
	options := newAwsRequestOptions(awsOpts...)

	baseOpts = append(
		baseOpts,
		WithHTTPClient(&awsClient{Service: service, AwsRequestOptions: *options}),
		WithToken("n/a"),
	)

	return NewClient(baseOpts...)
}

func (b *awsClient) Do(req *http.Request) (*http.Response, error) {
	isStream, err := b.setModelParams(req)
	if err != nil {
		return nil, err
	}

	err = signRequest(b.awsAccessKey, b.awsSecretKey, b.awsSessionToken, b.Service, b.awsRegion, req)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// map response to expected bedrock stream response
	if isStream && b.Service == "bedrock" {
		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// This is a naive implementation that assumes the responses are all in the same format
		var mappedResponseEvents []string
		events := strings.Split(string(resBody), "message")
		for _, event := range events {

			fmt.Println(event)
			// look for the payload
			idx := strings.Index(event, "\"bytes\":")
			if idx == -1 {
				continue
			}
			// get the payload and append it to the mapped response
			eventBody := strings.Split(event[idx:], "\"")[3]

			idx1 := strings.Index(eventBody, "stream-end")
			if idx1 != -1 {
				continue
			}

			// decode the payload
			decoded, err := base64.StdEncoding.DecodeString(eventBody)
			if err != nil {
				return nil, err
			}
			mappedResponseEvents = append(mappedResponseEvents, string(decoded))
		}

		resp.Body = io.NopCloser(strings.NewReader(strings.Join(mappedResponseEvents, "\n") + "\n"))
	}

	// parse response
	return resp, err
}

// modify the request to point to the bedrock model and remove the model from the request body
// handle stream param
func (b *awsClient) setModelParams(req *http.Request) (bool, error) {
	// try to parse the model from the request body
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return false, err
	}
	jsonBody := map[string]interface{}{}
	err = json.Unmarshal(reqBody, &jsonBody)
	if err != nil {
		return false, err
	}
	model, ok := jsonBody["model"].(string)
	if !ok {
		return false, errors.New("model not found in request body")
	}
	delete(jsonBody, "model")
	stream, ok := jsonBody["stream"].(bool)
	if !ok {
		stream = false
	} else if b.Service == "bedrock" {
		delete(jsonBody, "stream")
	}

	reqBody, err = json.Marshal(jsonBody)
	if err != nil {
		return false, err
	}

	req.URL, err = req.URL.Parse(getUrl(b.Service, b.awsRegion, model, stream))
	if err != nil {
		return false, err
	}

	req.Body = io.NopCloser(bytes.NewReader(reqBody))
	req.ContentLength = int64(len(reqBody))
	return stream, nil
}

// see https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func signRequest(accessID, secretKey, token, service, region string, req *http.Request) error {
	signer := v4.NewSigner()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	sha := sha256.New()
	_, err = sha.Write(bodyBytes)
	if err != nil {
		return err
	}
	payloadHash := hex.EncodeToString(sha.Sum(nil))

	err = signer.SignHTTP(
		req.Context(),
		aws.Credentials{AccessKeyID: accessID, SecretAccessKey: secretKey, SessionToken: token},
		req,
		payloadHash,
		service,
		region,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func getUrl(platform string, awsRegion string, model string, stream bool) string {
	endpoint := map[string]map[bool]string{
		"bedrock":   {true: "invoke-with-response-stream", false: "invoke"},
		"sagemaker": {true: "invocations-response-stream", false: "invocations"},
	}
	return fmt.Sprintf("https://%s-runtime.%s.amazonaws.com/model/%s/%s", platform, awsRegion, model, endpoint[platform][stream])
}

func NewBedrockClient(baseOpts []option.RequestOption, awsOpts []AwsRequestOption) *Client {
	return NewAwsClient(baseOpts, awsOpts, "bedrock")
}

func NewSagemakerClient(baseOpts []option.RequestOption, awsOpts []AwsRequestOption) *Client {
	return NewAwsClient(baseOpts, awsOpts, "sagemaker")
}
