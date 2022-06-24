package internal

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	ApiUrl     string
	HttpClient http.Client
}

func NewClient() *Client {
	return &Client{
		ApiUrl:     getApiUrlFromEnv(),
		HttpClient: http.Client{}, // TODO: think about params that may be needed in prod
	}
}

func getApiUrlFromEnv() string {
	const API_URL_ENV_VAR_NAME = "API_URL"
	apiUrl := os.Getenv(API_URL_ENV_VAR_NAME)
	validateUrl(apiUrl)
	return apiUrl
}

func validateUrl(urlFromEnv string) {
	_, err := url.Parse(urlFromEnv)
	if err != nil {
		log.Fatal(err)
	}
}
