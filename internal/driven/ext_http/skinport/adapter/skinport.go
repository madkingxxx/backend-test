package skinport

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://skinport.com/api"

	// average response time is ~1.2 to 2 seconds, so 5 seconds should be enough
	maxTimeoutSeconds = 5
	maxAttempts       = 3
)

type Sender struct {
	baseURL string
	client  http.Client
}

func New(baseURL string) *Sender {
	return &Sender{
		baseURL: baseURL,
		client: http.Client{
			Timeout: maxTimeoutSeconds * time.Second,
		},
	}
}

// isRetryableStatus checks if the HTTP status code indicates a retryable error.
// https://docs.skinport.com/introduction/status-codes
func isRetryableStatus(code int) bool {
	return code == http.StatusInternalServerError || code == http.StatusBadGateway ||
		code == http.StatusServiceUnavailable
}

// Do is a wrapper around http.Client.Do, for retrying requests if needed.
func (s *Sender) Do(req *http.Request) (*http.Response, error) {
	response, err := s.client.Do(req)

	// simple retry logic
	if isRetryableStatus(response.StatusCode) || err != nil {
		for i := 1; i <= maxAttempts; i++ {
			response, err = s.client.Do(req)
			if err != nil {
				continue
			}
			if !isRetryableStatus(response.StatusCode) && err == nil {
				return response, nil
			}
		}
	}
	return response, err
}
