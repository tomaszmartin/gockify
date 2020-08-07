package glockify

import (
	"net/http"
	"time"
)

const BaseURL = "https://api.clockify.me/api/v1"
const ReportURL = "https://reports.api.clockify.me/v1"

type Client struct {
	BaseURL    string
	ReportURL  string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL:   BaseURL,
		ReportURL: ReportURL,
		apiKey:    apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
