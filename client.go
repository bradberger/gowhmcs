package whmcs

const (
	// ErrClientNotFound is the current WHMCS response when the client does not exist.
	ErrClientNotFound string = "Client Not Found"
)

// NewClient contains the basic structure of data for creating a new client.
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
	Currency       string            `json:"currency"`
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

// AddClientResult is the WHMCS response when adding a client.
type AddClientResult struct {
	ClientID int64  `json:"clientid"`
	Result   string `json:"result"`
	Message  string `json:"message"`
}

// ClientDetailsReq is the struct of parameters available to retrieve client details.
type ClientDetailsReq struct {
	ClientID string `json:"clientid,omitempty"`
	Email    string `json:"email,omitempty"`
	Stats    bool   `json:"stats,omitempty"`
}

func (c *NewClient) Error() error {
	return nil
}
