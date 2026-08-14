package tests

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	core "github.com/cohere-ai/cohere-go/v2/core"
	option "github.com/cohere-ai/cohere-go/v2/option"
	"github.com/stretchr/testify/require"
)

type recordingHTTPClient struct {
	header http.Header
	calls  int
}

func (r *recordingHTTPClient) Do(req *http.Request) (*http.Response, error) {
	r.calls++
	r.header = req.Header.Clone()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"text":"hi"}`)),
	}, nil
}

// chatCapturing issues a request through client.NewClient with a recording HTTP client installed
// first, and returns the recorder so callers can assert on both the header and whether the
// caller's client was actually used.
func chatCapturing(t *testing.T, opts ...option.RequestOption) *recordingHTTPClient {
	t.Helper()
	recorder := &recordingHTTPClient{}
	co := client.NewClient(append([]option.RequestOption{option.WithHTTPClient(recorder)}, opts...)...)
	_, err := co.Chat(context.TODO(), &cohere.ChatRequest{Message: "hi"})
	require.NoError(t, err)
	return recorder
}

func chat(t *testing.T, opts ...option.RequestOption) http.Header {
	t.Helper()
	recorder := chatCapturing(t, opts...)
	require.Equal(t, 1, recorder.calls, "the caller's HTTP client must receive the request")
	return recorder.header
}

// An explicitly empty token means "send no Authorization header", and is distinct from never
// setting a token at all. The environment variable cannot reintroduce the header.
func TestEmptyTokenOmitsAuthorizationHeader(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	require.Empty(t, chat(t, option.WithToken("")).Get("Authorization"))
}

// The counterpart of the test above: omitting the option entirely still falls back to CO_API_KEY.
// These two together are what make WithToken("") meaningful rather than a no-op.
func TestOmittingTheTokenStillFallsBackToEnvironmentVariable(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	require.Equal(t, "Bearer env-token", chat(t).Get("Authorization"))
}

func TestEmptyTokenOmitsAuthorizationHeaderWhenEnvironmentIsUnset(t *testing.T) {
	t.Setenv("CO_API_KEY", "")
	require.Empty(t, chat(t, option.WithToken("")).Get("Authorization"))
}

// An empty token wins regardless of where it appears in the option list.
func TestEmptyTokenIsOrderIndependent(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	require.Empty(t, chat(t, option.WithToken("real-token"), option.WithToken("")).Get("Authorization"))
	require.Empty(t, chat(t, option.WithToken(""), option.WithToken("real-token")).Get("Authorization"))
}

// Suppressing auth must not swap out a caller's HTTP client, which would silently discard their
// proxy, mTLS config, timeouts and instrumentation.
func TestEmptyTokenPreservesCustomHTTPClient(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	recorder := chatCapturing(t, option.WithToken(""))
	require.Equal(t, 1, recorder.calls, "the caller's HTTP client must receive the request")
	require.Empty(t, recorder.header.Get("Authorization"))
}

func TestTokenIsSentWhenProvided(t *testing.T) {
	require.Equal(t, "Bearer some-token", chat(t, option.WithToken("some-token")).Get("Authorization"))
}

// Per-request suppression works the same as client-level: an empty token on a single call
// removes the Authorization header the client-scoped options had already set.
func TestPerRequestEmptyTokenSuppressesClientLevelAuth(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	recorder := &recordingHTTPClient{}
	co := client.NewClient(option.WithHTTPClient(recorder), option.WithToken("real-token"))
	_, err := co.Chat(context.TODO(), &cohere.ChatRequest{Message: "hi"}, option.WithToken(""))
	require.NoError(t, err)
	require.Empty(t, recorder.header.Get("Authorization"))
	require.Equal(t, 1, recorder.calls, "the caller's HTTP client must receive the request")
}

// Suppressing on one call must not affect the next one on the same client.
func TestPerRequestEmptyTokenDoesNotLeakIntoLaterRequests(t *testing.T) {
	recorder := &recordingHTTPClient{}
	co := client.NewClient(option.WithHTTPClient(recorder), option.WithToken("real-token"))

	_, err := co.Chat(context.TODO(), &cohere.ChatRequest{Message: "hi"}, option.WithToken(""))
	require.NoError(t, err)
	require.Empty(t, recorder.header.Get("Authorization"))

	_, err = co.Chat(context.TODO(), &cohere.ChatRequest{Message: "hi"})
	require.NoError(t, err)
	require.Equal(t, "Bearer real-token", recorder.header.Get("Authorization"))
}

// The suppression marker is an internal implementation detail and must never reach the wire, on
// any path. A leak would be worse than the bug this feature fixes.
func TestNoAuthMarkerNeverReachesTheWire(t *testing.T) {
	t.Setenv("CO_API_KEY", "env-token")
	for _, tt := range []struct {
		name       string
		clientOpts []option.RequestOption
		callOpts   []option.RequestOption
	}{
		{name: "client-level empty token", clientOpts: []option.RequestOption{option.WithToken("")}},
		{name: "per-request empty token", callOpts: []option.RequestOption{option.WithToken("")}},
		{
			name:       "both levels empty",
			clientOpts: []option.RequestOption{option.WithToken("")},
			callOpts:   []option.RequestOption{option.WithToken("")},
		},
		{name: "no suppression at all"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			recorder := &recordingHTTPClient{}
			opts := append([]option.RequestOption{option.WithHTTPClient(recorder)}, tt.clientOpts...)
			co := client.NewClient(opts...)
			_, err := co.Chat(context.TODO(), &cohere.ChatRequest{Message: "hi"}, tt.callOpts...)
			require.NoError(t, err)
			for key, values := range recorder.header {
				for _, value := range values {
					require.NotContains(t, value, core.NoAuthHeaderMarker,
						"marker leaked into header %q", key)
					require.NotContains(t, value, "fern-no-auth",
						"marker leaked into header %q", key)
				}
			}
		})
	}
}
