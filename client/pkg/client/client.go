package client

import (
	"net/http"
	"net/url"
)

type Client struct {
	apiUrl     *url.URL
	httpClient *http.Client

	AccountApi *accountsApiService
}

func (c Client) endpoint() *url.URL {
	return c.apiUrl
}

func NewClient(apiUrl *url.URL) *Client {
	c := &Client{
		apiUrl:     apiUrl,
		httpClient: &http.Client{},
	}
	c.AccountApi = newAccountsApiService(c)

	return c
}
