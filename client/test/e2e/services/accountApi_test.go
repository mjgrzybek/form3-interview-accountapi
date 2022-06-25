//go:build e2e

package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("create unique entry", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(accountRequestData)

		assert.Equal(t, ResponsesData["create"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("create entry already existing", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")
	})
}

func TestAccountApi_Fetch(t *testing.T) {
	t.Run("fetch existing entry", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["fetch"]
		accountResponseData, err := svc.Fetch(accountRequestData)
		assert.Equal(t, ResponsesData["fetch"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("fetch non-existing entry", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountRequestData := RequestsData["fetch"]
		uuid, err := uuid.NewUUID()
		if err != nil {
			t.Skipf("Unable to create new UUID %v", err)
		}
		accountRequestData.ID = uuid.String()

		accountResponseData, err := svc.Fetch(accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "record "+accountRequestData.ID+" does not exist")
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
