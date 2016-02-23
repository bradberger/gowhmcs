package whmcs

import (
    "encoding/json"
)

type CreateInvoiceRequest struct {

    UserID int64 `json:"userid,string"`
    Date InvoiceDate `json:"date"`
    DueDate InvoiceDate `json:"duedate"`
    PaymentMethod string `json:"paymentmethod"`
    ItemDescription1 string `json:"itemdescription1"`
    ItemAmount1 float64 `json:"itemamount1,string"`
    ItemTaxed1 int64 `json:"itemtaxed1,string"`

    // Optional attributes
    TaxRate int64 `json:"taxrate,string,omitempty"`
    TaxRate2 int64 `json:"taxrate2,string,omitempty"`
    Notes string `json:"notes,omitempty"`
    SendInvoice bool `json:"sendinvoice,string,omitempty"`
    AutoApplyCredit bool `json:"autoapplycredit,string,omitempty"`
    ItemDescription2 string `json:"itemdescription2,omitempty"`
    ItemAmount2 float64 `json:"itemamount2,string,omitempty"`
    ItemTaxed2 int64 `json:"itemtaxed2,string,omitempty"`

}

func (r *CreateInvoiceRequest) Error() error {
    return nil
}

type CreateInvoiceResponse struct {
    Result string `json:"result"`
    InvoiceID int64 `json:"invoiceid"`
}

type UpdateInvoiceResponse struct {
    Result string `json:"result"`
    InvoiceID int64 `json:"invoiceid,string"`
}

// CreateInvoice creates a new invoice with the given paramaters in `r`.
func (a *API) CreateInvoice(i *CreateInvoiceRequest) (r *CreateInvoiceResponse, err error) {

    err = i.Error()
    if err != nil {
        return
    }

    body, err := a.Do("createinvoice", i)
    if err != nil {
        return
    }

    r = &CreateInvoiceResponse{}
    err = json.Unmarshal(body, r)
    return

}

type UpdateInvoiceRequest struct {

    InvoiceID int64 `json:"invoiceid,string"`

    // Optional attributes
    Status string `json:"status,omitempty"`
    Date InvoiceDate `json:"date,omitempty"`
    DueDate InvoiceDate `json:"duedate,omitempty"`
    PaymentMethod string `json:"paymentmethod,omitempty"`
}

func (r *UpdateInvoiceRequest) Error() error {
    return nil
}

// UpdateInvoice updates the invoice with the given parameters of `r`.
func (a *API) UpdateInvoice(i *UpdateInvoiceRequest) (r *UpdateInvoiceResponse, err error) {

    err = i.Error()
    if err != nil {
        return
    }

    body, err := a.Do("updateinvoice", i)
    if err != nil {
        return
    }

    r = &UpdateInvoiceResponse{}
    err = json.Unmarshal(body, r)
    return

}
