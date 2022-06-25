package models

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"` // array [3] of string [140]
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	NameMatchingStatus      string   `json:"name_matching_status,omitempty"` // added
	Status                  *string  `json:"status,omitempty"`
	StatusReason            string   `json:"status_reason,omitempty"` // added
	Switched                *bool    `json:"switched,omitempty"`
	UserDefinedData         string   `json:"user_defined_data,omitempty"`    /// added
	ValidationType          string   `json:"validation_type,omitempty"`      // added
	ReferenceMask           string   `json:"reference_mask,omitempty"`       // added
	AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"` // added
}

type AccountDataRequest struct {
	Data *AccountData `json:"data,omitempty"`
}

type AccountDataResponse struct {
	Data *AccountData `json:"data,omitempty"`
}
