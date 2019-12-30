package whmcs

import (
	"encoding/json"
	"fmt"
)

// API is the main structure from which API calls are sent.
type API struct {
	Endpoint, Username, Password string
}

// AddTransaction creates the given transaction, returing error if it fails.
func (a *API) AddTransaction(t *Transaction) error {

	if err := t.Error(); err != nil {
		return fmt.Errorf("gowhmcs addtranscation error: %v", err)
	}

	if _, err := a.Do("addtransaction", &t); err != nil {
		return fmt.Errorf("gowhmcs addtranscation error: %v", err)
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
		err = fmt.Errorf("gowhmcs updateinvoice error: %v", err)
		return
	}

	r = &UpdateInvoiceResponse{}
	if err = json.Unmarshal(body, r); err != nil {
		err = fmt.Errorf("gowhmcs updateinvoice error: %v", err)
	}

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
		err = fmt.Errorf("gowhmcs updateexistingclient error: %v", err)
		return
	}

	body, err := a.Do("updateclient", &c)
	if err != nil {
		err = fmt.Errorf("gowhmcs updateexistingclient error: %v", err)
		return
	}

	r = &UpdateClientResult{}
	if err = json.Unmarshal(body, r); err != nil {
		err = fmt.Errorf("gowhmcs updateexistingclient error: %v", err)
	}
	return

}

// AcceptOrder accepts the WHMCS order with the matching parameters in `o`
func (a *API) AcceptOrder(o *AcceptOrderRequest) (err error) {

	err = o.Error()
	if err != nil {
		err = fmt.Errorf("gowhmcs acceptorder error: %v", err)
		return
	}

	if _, err = a.Do("acceptorder", o); err != nil {
		err = fmt.Errorf("gowhmcs acceptorder error: %v", err)
	}
	return

}

// AddOrder calls the "addorder" action of the WHMCS API.
func (a *API) AddOrder(o *Order) (r *OrderResponse, err error) {

	err = o.Error()
	if err != nil {
		err = fmt.Errorf("gowhmcs addorder error: %v", err)
		return
	}

	body, err := a.Do("addorder", o)
	if err != nil {
		err = fmt.Errorf("gowhmcs addorder error: %v (%s)", err, string(body))
		return
	}

	if err = json.Unmarshal(body, &r); err != nil {
		err = fmt.Errorf("gowhmcs addorder error: %s", string(body))
		return
	}
	return

}

// Whois returns the WHOIS data for the given domain.
func (a *API) Whois(domain string) (w WhoisInfo, err error) {

	body, err := a.Do("domainwhois", &WhoisRequest{"domainwhois", domain})
	if err != nil {
		err = fmt.Errorf("gowhmcs whois error: %v", err)
		return
	}

	if err = json.Unmarshal(body, &w); err != nil {
		err = fmt.Errorf("gowhmcs whois error: %v", err)
	}
	return

}

// ClientExists returns true if the client already exists.
func (a *API) ClientExists(email string) (bool, error) {
	c := ClientDetailsReq{Email: email}
	if _, err := a.Do("getclientsdetails", &c); err != nil {
		if err.Error() == ErrClientNotFound {
			return false, nil
		}
		return false, fmt.Errorf("gowhmcs clientexists error: %v", err)
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
		err = fmt.Errorf("gowhmcs addclient error: %v", err)
		return
	}

	r = &AddClientResult{}
	if err = json.Unmarshal(body, &r); err != nil {
		err = fmt.Errorf("gowhmcs addclient error: %v", err)
		return
	}

	return
}

func (a *API) ValidateLogin(v *ValidateLogin) (r *ValidateLoginResult, err error) {
	err = v.Error()
	if err != nil {
		return
	}

	body, err := a.Do("validatelogin", &v)
	if err != nil {
		err = fmt.Errorf("gowhmcs validatelogin error: %v", err)
		return
	}

	r = &ValidateLoginResult{}
	if err = json.Unmarshal(body, &r); err != nil {
		err = fmt.Errorf("gowhmcs validatelogin error: %v", err)
	}

	return
}

func (a *API) GetClientsProducts(v *GetClientsProducts) (r *GetClientsProductsResult, err error) {
    err = v.Error()
    if err != nil {
        return
    }
    body, err := a.Do("GetClientsProducts", &v)
    if err != nil {
        err = fmt.Errorf("%v", err)
        return
    }
    r = &GetClientsProductsResult{}
    if err = json.Unmarshal(body, &r); err != nil {
        err = fmt.Errorf("%v", err)
    }
    return
}

func (a *API) DecryptPassword(v *DecryptPassword) (r *DecryptPasswordResult, err error) {
    err = v.Error()
    if err != nil {
        return
    }
    body, err := a.Do("DecryptPassword", &v)
    if err != nil {
        err = fmt.Errorf("%v", err)
        return
    }
    r = &DecryptPasswordResult{}
    if err = json.Unmarshal(body, &r); err != nil {
        err = fmt.Errorf("%v", err)
    }
    return
}



func (a *API) UpdateClientProduct(p *ClientProduct) (r *UpdateClientProductResult, err error) {

	err = p.Error()
	if err != nil {
		return
	}

	body, err := a.Do("updateclientproduct", &p)
	if err != nil {
		err = fmt.Errorf("gowhmcs updateclientproduct error: %v", err)
		return
	}

	r = &UpdateClientProductResult{}
	if err = json.Unmarshal(body, r); err != nil {
		err = fmt.Errorf("gowhmcs updateclientproduct error: %v", err)
		return
	}

	return
}
