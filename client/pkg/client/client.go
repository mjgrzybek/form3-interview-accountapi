package client

import (
	"net/http"

	"github.com/mjgrzybek/form3-interview-accountapi/client/internal/services"
)

const (
	apiUrl = "https://api.form3.tech/v1"
)

type Client struct {
	httpClient http.Client

	AccountsApi services.AccountsApi
}
