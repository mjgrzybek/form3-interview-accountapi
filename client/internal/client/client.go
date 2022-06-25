package internal

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const ApiUrlEnvVarName = "API_URL"

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
	apiUrlEnvVarValue, ok := os.LookupEnv(ApiUrlEnvVarName)
	if !ok {
		log.Fatalf(`Environment variable "%s" not set`, ApiUrlEnvVarName)
	}

	apiUrl, err := url.Parse(apiUrlEnvVarValue)
	if err != nil {
		log.Fatal(err)
	}

	return apiUrl
}
