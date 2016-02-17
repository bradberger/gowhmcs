package whmcs

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
	Currency       int64             `json:"currency"`
	ClientIP       string            `json:"clientip"`
	Language       string            `json:"language"`
	GroupID        int64             `json:"groupid"`
	SecurityQID    int64             `json:"securityqid,string"`
	SecurityQans   string            `json:"securityqans"`
	Notes          string            `json:"notes"`
	CardNum        string            `json:"cardnum"`
	CardType       string            `json:"cardtype"`
	ExpDate        string            `json:"expdate"`
	StartDate      string            `json:"startdate"`
	IssueNumber    string            `json:"isseunumber"`
	CustomFields   map[string]string `json:"customfields"`
	NoEmail        bool              `json:"noemail"`
	SkipValidation bool              `json:"skipvalidation"`
}

func (c *Client) Error() error {
	return nil
}

type ClientDetailsReq struct {
	ClientID string `json:"clientid,omitempty"`
	Email    string `json:"email,omitempty"`
	Stats    bool   `json:"stats,omitempty"`
}

func (a *API) ClientExists(email string) (bool, error) {
	c := ClientDetailsReq{Email: email}
	if _, err := a.Do("getclientsdetails", &c); err != nil {
		return false, err
	}
	return true, nil
}
