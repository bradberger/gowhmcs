package whmcs

// CreateInvoiceRequest contains all available fields for creating a new invoice.
type CreateInvoiceRequest struct {
	UserID           int64       `json:"userid,string"`
	Date             InvoiceDate `json:"date"`
	DueDate          InvoiceDate `json:"duedate"`
	PaymentMethod    string      `json:"paymentmethod"`
	ItemDescription1 string      `json:"itemdescription1"`
	ItemAmount1      float64     `json:"itemamount1,string"`
	ItemTaxed1       int64       `json:"itemtaxed1,string"`

	// Optional attributes
	TaxRate          int64   `json:"taxrate,string,omitempty"`
	TaxRate2         int64   `json:"taxrate2,string,omitempty"`
	Notes            string  `json:"notes,omitempty"`
	SendInvoice      bool    `json:"sendinvoice,string,omitempty"`
	AutoApplyCredit  bool    `json:"autoapplycredit,string,omitempty"`
	ItemDescription2 string  `json:"itemdescription2,omitempty"`
	ItemAmount2      float64 `json:"itemamount2,string,omitempty"`
	ItemTaxed2       int64   `json:"itemtaxed2,string,omitempty"`
}

func (r *CreateInvoiceRequest) Error() error {
	return nil
}

// CreateInvoiceResponse is the WHMCS response when creating an invoice.
type CreateInvoiceResponse struct {
	Result    string `json:"result"`
	InvoiceID int64  `json:"invoiceid"`
}

// UpdateInvoiceResponse is the WHMCS response when updating an invoice.
type UpdateInvoiceResponse struct {
	Result    string `json:"result"`
	InvoiceID int64  `json:"invoiceid"`
}

// UpdateInvoiceRequest contains the parameters available to update an existing invoice.
type UpdateInvoiceRequest struct {
	InvoiceID int64 `json:"invoiceid,string"`

	// Optional attributes
	Status        string      `json:"status,omitempty"`
	Date          InvoiceDate `json:"date,omitempty"`
	DueDate       InvoiceDate `json:"duedate,omitempty"`
	PaymentMethod string      `json:"paymentmethod,omitempty"`
}

func (r *UpdateInvoiceRequest) Error() error {
	return nil
}
