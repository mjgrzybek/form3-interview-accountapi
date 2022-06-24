package internal

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	ApiUrl     *url.URL
	HttpClient http.Client
}

func NewClient() *Client {
	return &Client{
		ApiUrl:     getApiUrlFromEnv(),
		HttpClient: http.Client{}, // TODO: think about params that may be needed in prod
	}
}

func getApiUrlFromEnv() *url.URL {
	const API_URL_ENV_VAR_NAME = "API_URL"

	urlFromEnv := os.Getenv(API_URL_ENV_VAR_NAME)
	apiUrl, err := url.Parse(urlFromEnv)

	if err != nil {
		log.Fatal(err)
	}

	return apiUrl
}
