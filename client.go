package whmcs

import (
	"encoding/json"
	"log"
)

const (
	ErrClientNotFound string = "Client Not Found"
)

type Client struct {
	NewClient
	ExistingClient
}

type NewClient struct {

	// Required if adding.
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Address1    string `json:"address1"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phonenumber"`
	Password2   string `json:"password2"`

	// Optional
	CompanyName    string            `json:"companyname"`
	Address2       string            `json:"address2"`
	Currency       string             `json:"currency"`
	ClientIP       string            `json:"clientip"`
	Language       string            `json:"language"`
	GroupID        int64             `json:"groupid,string"`
	SecurityQID    int64             `json:"securityqid,string"`
	SecurityQans   string            `json:"securityqans"`
	Notes          string            `json:"notes"`
	CardNum        string            `json:"cardnum"`
	CardType       string            `json:"cardtype"`
	ExpDate        string            `json:"expdate"`
	StartDate      string            `json:"startdate"`
	IssueNumber    string            `json:"isseunumber"`
	CustomFields   map[string]string `json:"customfields"`
	NoEmail        bool              `json:"noemail,string"`
	SkipValidation bool              `json:"skipvalidation,string"`
}

func (c *NewClient) Error() error {
	// @TODO Add error checks here.
	return nil
}

type AddClientResult struct {
	ClientID int64 `json:"clientid"`
	Result string `json:"result"`
	Message string `json:"message"`
}

type ClientDetailsReq struct {
	ClientID string `json:"clientid,omitempty"`
	Email    string `json:"email,omitempty"`
	Stats    bool   `json:"stats,omitempty"`
}

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
	log.Printf("Addclient: %+v %v\n", r, err)
	return

}
