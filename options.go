package pavlok

import (
	"net/http"
)

type OptionsFunc func(client *Client) *Client

func WithHTTPClient(httpClient *http.Client) OptionsFunc {
	return func(client *Client) *Client {
		client.httpClient = httpClient
		return client
	}
}

func WithBaseURL(baseURL string) OptionsFunc {
	return func(client *Client) *Client {
		client.baseURL = baseURL
		return client
	}
}
