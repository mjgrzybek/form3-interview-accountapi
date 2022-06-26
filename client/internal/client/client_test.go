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

		client, err := NewClient()
		assert.NoError(t, err)
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

		client, err := NewClient()
		assert.NoError(t, err)
		assert.NotNil(t, client.HttpClient)
		assert.Equal(t, parsedUrl, client.ApiUrl)
	})
	t.Run("url is invalid", func(t *testing.T) {
		const apiUrl = string(byte(0x7f))

		os.Setenv(ApiUrlEnvVarName, apiUrl)
		defer os.Unsetenv(ApiUrlEnvVarName)

		client, err := NewClient()
		assert.EqualError(t, err, "parse \"\\u007f\": net/url: invalid control character in URL")
		assert.Nil(t, client)
	})
	t.Run("url is not set", func(t *testing.T) {
		os.Unsetenv(ApiUrlEnvVarName)
		client, err := NewClient()

		assert.EqualError(t, err, "Environment variable \"API_URL\" not set")
		assert.Nil(t, client)
	})
}
