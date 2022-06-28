package client

import (
	"net/http"
	"net/url"
)

type Client struct {
	ApiUrl     *url.URL
	HttpClient *http.Client

	AccountApi *accountsApiService
}

func (c Client) endpoint() *url.URL {
	return c.ApiUrl
}

func NewClient(apiUrl *url.URL) *Client {
	c := &Client{
		ApiUrl:     apiUrl,
		HttpClient: &http.Client{},
	}
	c.AccountApi = newAccountsApiService(c)

	return c
}
