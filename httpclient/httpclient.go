package httpclient

import (
	"io"
	"net/http"
	"time"
)

// HTTPClient interface for http communication
type HTTPClient interface {
	SendGetRequest(uri string, headers map[string]string) ([]byte, error)
}

// Impl implements httpClient interface
type Impl struct {
}

func getClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	return &http.Client{Transport: tr}
}

// SendGetRequest will send a "GET" request to the uri with the optional additional headers. Will return a byte array of the body or error.
func (h *Impl) SendGetRequest(uri string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	for k, value := range headers {
		req.Header.Set(k, value)
	}
	client := getClient()
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
