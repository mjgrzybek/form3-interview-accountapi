package internal

import (
	"net/http"
)

type Client struct {
	HttpClient http.Client
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.Client{}, // TODO: think about params that may be needed in prod
	}
}
