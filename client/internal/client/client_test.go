package internal

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	preparedUrl, _ := url.Parse("http://asd")

	type args struct {
		apiUrl *url.URL
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "simple",
			args: args{
				apiUrl: preparedUrl,
			},
			want: &Client{
				ApiUrl:     preparedUrl,
				HttpClient: http.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
