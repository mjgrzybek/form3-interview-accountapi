//go:build e2e

package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(accountRequestData)

		assert.Equal(t, ResponsesData["create"], accountResponseData)
		assert.NoError(t, err)
	})
}

func TestAccountApi_Delete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["delete"]
		err := svc.Delete(accountRequestData)
		assert.NoError(t, err)
	})
}
