package whmcs

import (
	"encoding/json"
)

// API is the main structure from which API calls are sent.
type API struct {
	Endpoint, Username, Password string
}

// AddTransaction creates the given transaction, returing error if it fails.
func (a *API) AddTransaction(t *Transaction) error {

	if err := t.Error(); err != nil {
		return err
	}

	if _, err := a.Do("addtransaction", &t); err != nil {
		return err
	}

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

// UpdateExistingClient updates an existing client.
func (a *API) UpdateExistingClient(c *ExistingClient) (r *UpdateClientResult, err error) {

	err = c.Error()
	if err != nil {
		return
	}

	body, err := a.Do("updateclient", &c)
	if err != nil {
		return
	}

	r = &UpdateClientResult{}
	err = json.Unmarshal(body, r)
	return

}

// AcceptOrder accepts the WHMCS order with the matching parameters in `o`
func (a *API) AcceptOrder(o *AcceptOrderRequest) (err error) {

	err = o.Error()
	if err != nil {
		return
	}

	_, err = a.Do("acceptorder", o)
	return

}

// AddOrder calls the "addorder" action of the WHMCS API.
func (a *API) AddOrder(o *Order) (r *OrderResponse, err error) {

	err = o.Error()
	if err != nil {
		return
	}

	body, err := a.Do("addorder", o)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &r)
	return

}

// Whois returns the WHOIS data for the given domain.
func (a *API) Whois(domain string) (w WhoisInfo, err error) {

	body, err := a.Do("domainwhois", &WhoisRequest{"domainwhois", domain})
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &w)
	return

}

// ClientExists returns true if the client already exists.
func (a *API) ClientExists(email string) (bool, error) {
	c := ClientDetailsReq{Email: email}
	if _, err := a.Do("getclientsdetails", &c); err != nil {
		if err.Error() == ErrClientNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AddClient creates a new WHMCS client.
func (a *API) AddClient(c *NewClient) (r *AddClientResult, err error) {

	err = c.Error()
	if err != nil {
		return
	}

	body, err := a.Do("addclient", &c)
	if err != nil {
		return
	}

	r = &AddClientResult{}
	err = json.Unmarshal(body, &r)
	return
}

func (a *API) UpdateClientProduct (p *ClientProduct) (r *UpdateClientProductResult, err error) {

	err = p.Error()
	if err != nil {
		return
	}

	body, err := a.Do("updateclientproduct", &p)
	if err != nil {
		return
	}

	r = &UpdateClientProductResult{}
	err = json.Unmarshal(body, r)
	return
}
