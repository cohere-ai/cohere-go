package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	core "github.com/cohere-ai/cohere-go/v2/core"
)

// HTTPClient is an interface for a subset of the *http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ResolveBaseURL resolves the base URL from the given arguments,
// preferring the first non-empty value.
func ResolveBaseURL(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

// EncodeURL encodes the given arguments into the URL, escaping
// values as needed. Pointer arguments are dereferenced before processing.
func EncodeURL(urlFormat string, args ...interface{}) string {
	escapedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		// Dereference the argument if it's a pointer
		value := dereferenceArg(arg)
		escapedArgs = append(escapedArgs, url.PathEscape(fmt.Sprintf("%v", value)))
	}
	return fmt.Sprintf(urlFormat, escapedArgs...)
}

// dereferenceArg dereferences a pointer argument if necessary, returning the underlying value.
// If the argument is not a pointer or is nil, it returns the argument as-is.
func dereferenceArg(arg interface{}) interface{} {
	if arg == nil {
		return arg
	}

	v := reflect.ValueOf(arg)

	// Keep dereferencing until we get to a non-pointer value or hit nil
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	return v.Interface()
}

// MergeHeaders merges the given headers together, where the right
// takes precedence over the left.
//
// FERN: hand-maintained. Headers carrying core.NoAuthHeaderMarker are deleted rather than
// merged, which is how option.WithToken("") suppresses an 'Authorization' header that the
// client-scoped options already set. The marker is stripped from both sides, so it can never
// reach the wire. Every request path funnels through this function. See .fernignore.
func MergeHeaders(left, right http.Header) http.Header {
	for key, values := range right {
		if len(values) > 1 {
			left[key] = values
			continue
		}
		value := right.Get(key)
		if value == core.NoAuthHeaderMarker {
			left.Del(key)
			continue
		}
		if value != "" {
			left.Set(key, value)
		}
	}
	// The marker may also arrive on the left, from client-scoped options.
	for key := range left {
		if left.Get(key) == core.NoAuthHeaderMarker {
			left.Del(key)
		}
	}
	return left
}
