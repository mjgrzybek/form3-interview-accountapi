package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("url is set", func(t *testing.T) {
		const url = "boo"

		os.Setenv(ApiUrlEnvVarName, url)
		defer os.Unsetenv(ApiUrlEnvVarName)

		client := NewClient()
		assert.NotNil(t, client.HttpClient)
		assert.Equal(t, url, client.ApiUrl)
	})
	t.Run("url is set to an empty string", func(t *testing.T) {
		const url = ""

		os.Setenv(ApiUrlEnvVarName, url)
		defer os.Unsetenv(ApiUrlEnvVarName)

		client := NewClient()
		assert.NotNil(t, client.HttpClient)
		assert.Equal(t, url, client.ApiUrl)
	})
	t.Run("url is not set", func(t *testing.T) {
		// os.Exit exits unit test
		//
		//client := NewClient()
		//assert.NotNil(t, client.HttpClient)
		//assert.Equal(t, "", client.ApiUrl)
	})
}
