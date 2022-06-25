package internal

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("url is set", func(t *testing.T) {
		const apiUrl = "boo"

		os.Setenv(ApiUrlEnvVarName, apiUrl)
		defer os.Unsetenv(ApiUrlEnvVarName)

		parsedUrl, err := url.Parse(apiUrl)
		if err != nil {
			t.Skipf("Couldn't parse url")
		}

		client := NewClient()
		assert.NotNil(t, client.HttpClient)
		assert.Equal(t, parsedUrl, client.ApiUrl)
	})
	t.Run("url is set to an empty string", func(t *testing.T) {
		const apiUrl = ""

		os.Setenv(ApiUrlEnvVarName, apiUrl)
		defer os.Unsetenv(ApiUrlEnvVarName)

		parsedUrl, err := url.Parse(apiUrl)
		if err != nil {
			t.Skipf("Couldn't parse url")
		}

		client := NewClient()
		assert.NotNil(t, client.HttpClient)
		assert.Equal(t, parsedUrl, client.ApiUrl)
	})
	t.Run("url is not set", func(t *testing.T) {
		// os.Exit exits unit test
		//
		//client := NewClient()
		//assert.NotNil(t, client.HttpClient)
		//assert.Equal(t, "", client.ApiUrl)
	})
}
