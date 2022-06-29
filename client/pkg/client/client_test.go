package client

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	utils "github.com/mjgrzybek/form3-interview-accountapi/client/internal/utils"
)

func TestClient_endpoint(t *testing.T) {
	preparedUrl, _ := url.Parse("http://asd")

	type fields struct {
		ApiUrl     *url.URL
		HttpClient *http.Client
		AccountApi *accountsApiService
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		{
			name: "simple",
			fields: fields{
				ApiUrl:     preparedUrl,
				HttpClient: &http.Client{},
				AccountApi: &accountsApiService{
					Client:   nil,
					Endpoint: utils.JoinPathUrl(*preparedUrl, "organisation", "accounts"),
				},
			},
			want: preparedUrl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				apiUrl:     tt.fields.ApiUrl,
				httpClient: tt.fields.HttpClient,
				AccountApi: tt.fields.AccountApi,
			}
			if got := c.endpoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		apiUrl *url.URL
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
