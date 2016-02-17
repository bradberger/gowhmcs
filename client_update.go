package whmcs

import (
    "errors"
)

var (
    ErrNoClientDetails = errors.New("Client must have either ClientEmail or ClientID")
)

type ExistingClient struct {
	ClientID    string `"json:clientid,omitempty"`
	ClientEmail string `json:"clientemail,omitempty"`

	Firstname           string       `json:"firstname,omitempty"`
	Lastname            string       `json:"lastname,omitempty"`
	CompanyName         string       `json:"companyname,omitempty"`
	Email               string       `json:"email,omitempty"`
	Address1            string       `json:"address1,omitempty"`
	Address2            string       `json:"address2,omitempty"`
	City                string       `json:"city,omitempty"`              //
	State               string       `json:"state,omitempty"`             //
	Postcode            string       `json:"postcode,omitempty"`          //
	Country             string       `json:"country,omitempty"`           // - two letter ISO country code
	PhoneNumber         string       `json:"phonenumber,omitempty"`       //
	Password2           string       `json:"password2,omitempty"`         //
	Credit              float64      `json:"credit,string,omitempty"`     // - credit balance
	TaxExempt           bool         `json:"taxexempt,omitempty"`         // - true to enable
	Notes               string       `json:"notes,omitempty"`             //
	CardType            string       `json:"cardtype,omitempty"`          // - visa, mastercard, etc...
	CardNum             string       `json:"cardnum,omitempty"`           // - cc number
	ExpDate             Date         `json:"expdate,omitempty"`           // - cc expiry date
	StartDate           Date         `json:"startdate,omitempty"`         // - cc start date
	IssueNumber         string       `json:"issuenumber,omitempty"`       // - cc issue number
	BankName            string       `json:"bankname,omitempty"`          // - for use with direct debit gateway
	BankType            string       `json:"banktype,omitempty"`          // - for use with direct debit gateway
	BankCode            string       `json:"bankcode,omitempty"`          // - for use with direct debit gateway
	BankAcct            string       `json:"bankacct,omitempty"`          // - for use with direct debit gateway
	Language            string       `json:"language,omitempty"`          // - default language
	ClearCreditCard     bool         `json:"clearcreditcard,omitempty"`   // - set to true to remove stored card data (V5.3.7+ only)
	PaymentMethod       string       `json:"paymentmethod,omitempty"`     // - paypal, authorize, etc...
	CustomFiels         CustomFields `json:"customfields,omitempty"`      // - a base64 encoded serialized array of custom field values
	Status              string       `json:"status,omitempty"`            // - active or inactive
	LateFeeOverride     bool         `json:"latefeeoveride,omitempty"`    // - true/false
	OverrideDueInvoices bool         `json:"overideduenotices,omitempty"` // - true/false
	SeparateInvoices    bool         `json:"separateinvoices,omitempty"`  // - true/false
	DisableAutoCC       bool         `json:"disableautocc,omitempty"`     // - true/false

}

func (c *ExistingClient) Error() error {
    if c.ClientID == "" && c.ClientEmail == "" {
        return ErrNoClientDetails
    }
	return nil
}

func (c *ExistingClient) ByEmail(email string) {
	c.ClientEmail = email
}

func (c *ExistingClient) ByID(id string) {
	c.ClientID = id
}
