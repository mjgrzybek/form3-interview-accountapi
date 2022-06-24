//go:build e2e

package services

import (
	"io"
	"testing"

	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()

		accountData := models.AccountDataRequest{
			Data: &models.AccountData{
				Attributes: &models.AccountAttributes{
					AccountClassification:   nil,
					AccountMatchingOptOut:   nil,
					AccountNumber:           "",
					AlternativeNames:        nil,
					BankID:                  "",
					BankIDCode:              "",
					BaseCurrency:            "",
					Bic:                     "",
					Country:                 nil,
					Iban:                    "",
					JointAccount:            nil,
					Name:                    nil,
					SecondaryIdentification: "",
					Status:                  nil,
					Switched:                nil,
				},
				ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
				Type:           "accounts",
			},
		}
		response, err := svc.Create(accountData)

		t.Log(response.StatusCode, reader2str(response.Body))
		assert.NoError(t, err)
	})
}

func reader2str(reader io.Reader) string {
	all, _ := io.ReadAll(reader)
	return string(all)
}
