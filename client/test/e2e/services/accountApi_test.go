//go:build e2e

package services

import (
	"testing"

	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"

	client "github.com/mjgrzybek/form3-interview-accountapi/client/pkg/services"
)

func TestAccountApi_Create(t *testing.T) {
	t.Run("", func(t *testing.T) {
		svc := client.NewAccountsApiService()
		
		err := svc.Create(&models.AccountData{
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
			ID:             "",
			OrganisationID: "",
			Type:           "",
			Version:        nil,
		})

		assert.NoError(t, err)
	})
}
