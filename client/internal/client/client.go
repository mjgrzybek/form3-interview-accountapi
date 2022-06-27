package internal

import (
	"errors"
	"net/http"
	"net/url"
	"os"
)

const ApiUrlEnvVarName = "API_URL"

type Client struct {
	ApiUrl     *url.URL
	HttpClient http.Client
}

func NewClient() (*Client, error) {
	apiUrl, err := getApiUrlFromEnv() // TODO: switch to parameter instead of envvar
	if err != nil {
		return nil, err
	}

	return &Client{
		ApiUrl:     apiUrl,
		HttpClient: http.Client{},
	}, nil
}

func getApiUrlFromEnv() (*url.URL, error) {
	apiUrlEnvVarValue, ok := os.LookupEnv(ApiUrlEnvVarName)
	if !ok {
		return nil, errors.New("Environment variable \"" + ApiUrlEnvVarName + "\" not set")
	}

	apiUrl, err := url.Parse(apiUrlEnvVarValue)
	if err != nil {
		return nil, err
	}

	return apiUrl, nil
}
