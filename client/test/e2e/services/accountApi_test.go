//go:build e2e

package services

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountData := RequestsData["create"]
		httpResponse, err := svc.Create(accountData)
		assert.Equal(t, http.StatusCreated, httpResponse.StatusCode)

		var response models.AccountDataResponse
		json.Unmarshal(reader2bytes(httpResponse.Body), &response)
		assert.Equal(t, ResponsesData["create"], response.Data)
		assert.NoError(t, err)
	})
}

func TestAccountApi_Delete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountData := RequestsData["delete"]
		httpResponse, err := svc.Delete(accountData)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, httpResponse.StatusCode)
	})
}

func reader2str(reader io.Reader) string {
	all, _ := io.ReadAll(reader)
	return string(all)
}

func reader2bytes(reader io.Reader) []byte {
	all, _ := io.ReadAll(reader)
	return all
}
