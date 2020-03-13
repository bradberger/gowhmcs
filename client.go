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
	CompanyName    string            `json:"companyname,omitempty"`
	Address2       string            `json:"address2,omitempty"`
	Currency       string            `json:"currency,omitempty"`
	ClientIP       string            `json:"clientip,omitempty"`
	Language       string            `json:"language,omitempty"`
	GroupID        int64             `json:"groupid,string,omitempty"`
	SecurityQID    int64             `json:"securityqid,string,omitempty"`
	SecurityQans   string            `json:"securityqans,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	CardNum        string            `json:"cardnum,omitempty"`
	CardType       string            `json:"cardtype,omitempty"`
	ExpDate        string            `json:"expdate,omitempty"`
	StartDate      string            `json:"startdate,omitempty"`
	IssueNumber    string            `json:"isseunumber,omitempty"`
	CustomFields   map[string]string `json:"customfields,omitempty"`
	NoEmail        bool              `json:"noemail,string,omitempty"`
	SkipValidation bool              `json:"skipvalidation,string,omitempty"`
}

type ValidateLogin struct {
	Email     string `json:"email"`
	Password2 string `json:"password2"`
}

type DecryptPassword struct {
	Password2 string `json:"password2"`
}

type ResetPassword struct {
	Email string `json:"email"`
}

type TerminateService struct {
	Id int64 `json:"id"`
}

type CreateService struct {
	Id int64 `json:"serviceid"`
}

type OpenTicketRequest struct {
	DeptId   int64  `json:"deptid"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	ClientId int64  `json:"clientid"`
}

type GetClientsProducts struct {
	ServiceId int `json:"serviceid"`
}

type GetClientsProductsResult struct {
	Result       string `json:"result"`
	ClientID     int64  `json:"clientid"`
	ServiceID    string `json:"serviceid"`
	Pid          int64  `json:"pid"`
	Domain       string `json:"domain"`
	TotalResults int64  `json:"totalresults"`
	StartNumber  int64  `json:"startnumber"`
	NumReturned  int64  `json:"numreturned"`
}

// AddClientResult is the WHMCS response when adding a client.
type AddClientResult struct {
	ClientID int64  `json:"clientid"`
	Result   string `json:"result"`
	Message  string `json:"message"`
}

type ValidateLoginResult struct {
	UserID       int64  `json:"userid"`
	ContactID    int64  `json:"contactid"`
	PasswordHash string `json:"passwordhash"`
	Result       string `json:"result"`
	Message      string `json:"message"`
}

type DecryptPasswordResult struct {
	Result   string `json:"result"`
	Password string `json:"password"`
}

type ResetPasswordResult struct {
	Result string `json:"result"`
}

type TerminateServiceResult struct {
	Result string `json:"result"`
}

type CreateServiceResult struct {
	Result string `json:"result"`
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

func (v *ValidateLogin) Error() error {
	return nil
}

func (v *ResetPassword) Error() error {
	return nil
}

func (v *GetClientsProducts) Error() error {
	return nil
}

func (v *DecryptPassword) Error() error {
	return nil
}

func (v *TerminateService) Error() error {
	return nil
}
func (v *CreateService) Error() error {
	return nil
}

func (v *OpenTicketRequest) Error() error {
	return nil
}

type OpenTicketResponse struct {
	Result string `json:"result"`
}
