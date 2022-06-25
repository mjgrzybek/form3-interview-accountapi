//go:build e2e

package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("create unique entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService()
		assert.NoError(t, err)

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(accountRequestData)

		assert.Equal(t, ResponsesData["create"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("create entry already existing", func(t *testing.T) {
		svc, err := client.NewAccountsApiService()
		assert.NoError(t, err)

		accountRequestData := RequestsData["create"]
		accountResponseData, err := svc.Create(accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")
	})
}

func TestAccountApi_Fetch(t *testing.T) {
	t.Run("fetch existing entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService()
		assert.NoError(t, err)

		accountRequestData := RequestsData["fetch"]
		accountResponseData, err := svc.Fetch(accountRequestData)
		assert.Equal(t, ResponsesData["fetch"], accountResponseData)
		assert.NoError(t, err)
	})
	t.Run("fetch non-existing entry", func(t *testing.T) {
		svc, err := client.NewAccountsApiService()
		assert.NoError(t, err)

		accountRequestData := RequestsData["fetch"]
		err = setRandomUuid(t, accountRequestData)

		accountResponseData, err := svc.Fetch(accountRequestData)

		assert.Nil(t, accountResponseData)
		assert.EqualError(t, err, "record "+accountRequestData.ID+" does not exist")
	})
}

func setRandomUuid(t *testing.T, accountRequestData *models.AccountData) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Skipf("Unable to create new UUID %v", err)
	}
	accountRequestData.ID = uuid.String()
	return err
}

func TestAccountApi_Delete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc, err := client.NewAccountsApiService()
		assert.NoError(t, err)

		accountRequestData := RequestsData["delete"]
		err = svc.Delete(accountRequestData)
		assert.NoError(t, err)
	})
}
