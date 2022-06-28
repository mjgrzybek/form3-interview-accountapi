package client

import (
	"github.com/mjgrzybek/form3-interview-accountapi/client/internal/address"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
)

var CreateRequestsData = &models.AccountData{
	AccountIdVersion: models.AccountIdVersion{
		AccountId: models.AccountId{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
	},
	OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	Type:           "accounts",
	Attributes: models.AccountAttributes{
		BankID:              "400300",
		BankIDCode:          "GBDSC",
		BaseCurrency:        "GBP",
		Bic:                 "NWBKGB22",
		Country:             address.Of("GB"),
		ValidationType:      "card",
		ReferenceMask:       "############",
		AcceptanceQualifier: "same_day",
		Name:                []string{"Samantha Holder"},
		AlternativeNames:    []string{"Sam Holder"},
	},
}

var FetchRequestsData = &models.AccountIdVersion{
	AccountId: models.AccountId{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
}

var DeleteRequestData = &models.AccountIdVersion{
	AccountId: models.AccountId{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
	Version:   address.Of[int64](0),
}

var CreateResponseData = &models.AccountData{
	AccountIdVersion: models.AccountIdVersion{
		AccountId: models.AccountId{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
		Version:   address.Of[int64](0),
	},
	OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	Type:           "accounts",
	Attributes: models.AccountAttributes{
		AlternativeNames: []string{"Sam Holder"},
		BankID:           "400300",
		BankIDCode:       "GBDSC",
		BaseCurrency:     "GBP",
		Bic:              "NWBKGB22",
		Country:          address.Of("GB"),
		Name:             []string{"Samantha Holder"},
	},
}

var FetchResponseData = &models.AccountData{
	AccountIdVersion: models.AccountIdVersion{
		AccountId: models.AccountId{ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
		Version:   address.Of[int64](0),
	},
	OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	Type:           "accounts",
	Attributes: models.AccountAttributes{
		AlternativeNames: []string{"Sam Holder"},
		BankID:           "400300",
		BankIDCode:       "GBDSC",
		BaseCurrency:     "GBP",
		Bic:              "NWBKGB22",
		Country:          address.Of("GB"),
		Name:             []string{"Samantha Holder"},
	},
}
