package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoints(t *testing.T) {
	r := Setup()
	ts := httptest.NewServer(r)
	defer ts.Close()

	tests := []struct {
		endpoint   string
		statusCode int
	}{
		{
			endpoint:   "/v1/advice",
			statusCode: http.StatusOK,
		},
		{
			endpoint:   "/v1/chucknorris",
			statusCode: http.StatusOK,
		},
		{
			endpoint:   "/v1/dadjoke",
			statusCode: http.StatusOK,
		},
		{
			endpoint:   "/v1/random",
			statusCode: http.StatusOK,
		},
		{
			endpoint:   "/some/bad/endpoint",
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			res, err := http.Get(ts.URL + tt.endpoint)
			require.NoError(t, err, "There should not be an error in the server.")
			require.Equal(t, tt.statusCode, res.StatusCode, "Unexpected response status code")
		})
	}
}
