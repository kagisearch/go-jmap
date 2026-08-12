package jmap

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sessionErrorClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)
	return &Client{SessionEndpoint: ts.URL, HttpClient: ts.Client()}
}

// A non-JSON error response must still name the status code: telling a 401
// (rejected credentials) apart from a 503 (transient server failure) is the
// whole point of the error.
func TestAuthenticateSurfacesHTTPStatus(t *testing.T) {
	for _, tc := range []struct {
		status int
		want   string
	}{
		{http.StatusUnauthorized, "HTTP 401 Unauthorized"},
		{http.StatusServiceUnavailable, "HTTP 503 Service Unavailable"},
	} {
		c := sessionErrorClient(t, func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "nope", tc.status)
		})

		err := c.Authenticate()
		require.Error(t, err)
		assert.Contains(t, err.Error(), tc.want)
		assert.Contains(t, err.Error(), "jmap session")
	}
}

// A JSON problem-details body decodes to a *RequestError, which callers can
// match on, and whose message carries the status code and detail.
func TestAuthenticateSurfacesRequestError(t *testing.T) {
	c := sessionErrorClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{
			"type": "urn:ietf:params:jmap:error:limit",
			"status": 429,
			"detail": "too many concurrent requests",
			"limit": "concurrentRequests"
		}`))
	})

	err := c.Authenticate()
	require.Error(t, err)

	var reqErr *RequestError
	require.True(t, errors.As(err, &reqErr))
	assert.Equal(t, 429, reqErr.Status)
	assert.Equal(t, "HTTP 429: too many concurrent requests: concurrentRequests", reqErr.Error())
}

// Problem details arrive as "application/problem+json", and any JSON media type
// may carry parameters. Matching the header exactly would drop the body and
// leave nothing but the status line.
func TestAuthenticateRequestErrorContentTypes(t *testing.T) {
	for _, contentType := range []string{
		"application/json",
		"application/json; charset=utf-8",
		"application/problem+json",
		"application/problem+json; charset=utf-8",
	} {
		t.Run(contentType, func(t *testing.T) {
			c := sessionErrorClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", contentType)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{
					"type": "urn:ietf:params:jmap:error:limit",
					"status": 400,
					"detail": "The request exceeds the maximum number of concurrent requests.",
					"limit": "maxConcurrentRequests"
				}`))
			})

			err := c.Authenticate()
			require.Error(t, err)

			var reqErr *RequestError
			require.True(t, errors.As(err, &reqErr))
			assert.Equal(t, "urn:ietf:params:jmap:error:limit", reqErr.Type)
			assert.Contains(t, err.Error(), "The request exceeds the maximum number of concurrent requests.")
			assert.Contains(t, err.Error(), "maxConcurrentRequests")
		})
	}
}

// A JSON body that omits "status" still reports the response's status code.
func TestAuthenticateRequestErrorWithoutStatus(t *testing.T) {
	c := sessionErrorClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"detail": "account is disabled"}`))
	})

	err := c.Authenticate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP 403: account is disabled")
}

func TestAuthenticateSucceeds(t *testing.T) {
	c := sessionErrorClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(sessionBlob))
	})

	require.NoError(t, c.Authenticate())
	assert.Equal(t, "john@example.com", c.Session.Username)
}
