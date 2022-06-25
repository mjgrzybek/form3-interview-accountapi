package internal

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const ApiUrlEnvVarName = "API_URL"

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
	apiUrl, ok := os.LookupEnv(ApiUrlEnvVarName)
	if !ok {
		log.Fatalf(`Environment variable "%s" not set`, ApiUrlEnvVarName)
	}
	validateUrl(apiUrl)
	return apiUrl
}

func validateUrl(urlFromEnv string) {
	_, err := url.Parse(urlFromEnv)
	if err != nil {
		log.Fatal(err)
	}
}
