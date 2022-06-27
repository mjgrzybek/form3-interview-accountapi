package internal

import (
	"net/http"
	"net/url"
)

type Client struct {
	ApiUrl     *url.URL
	HttpClient http.Client
}

func NewClient(apiUrl *url.URL) *Client {
	return &Client{
		ApiUrl:     apiUrl,
		HttpClient: http.Client{},
	}
}
