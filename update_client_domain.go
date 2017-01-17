package whmcs

import (
	"encoding/json"
	"fmt"
	"time"
)

type ClientDomainUpdateReq struct {
	DomainID int64 `json:"domainid"`

	DNSManagement      bool    `json:"dnsmanagement,omitempty"`
	EmailForwarding    bool    `json:"emailforwarding,omitempty"`
	IDProtection       bool    `json:"idprotection,omitempty"`
	DoNotRenew         bool    `json:"donotrenew,omitempty"`
	Type               string  `json:"type,omitempty"`
	RegDate            string  `json:"regdate,omitempty"`     // Y-m-d
	NextDueDate        string  `json:"nextduedate,omitempty"` // Y-m-d
	ExpiryDate         string  `json:"expirydate,omitempty"`  // Y-m-d
	Domain             string  `json:"domain,omitempty"`
	FirstPaymentAmount float64 `json:"firstpaymentamount,omitempty"`
	Registrar          string  `json:"registrar,omitempty"`
	RegPeriod          int64   `json:"regperiod,omitempty"`
	PaymentMethod      string  `json:"paymentmethod,omitempty"`
	SubscriptionID     string  `json:"subscriptionid,omitempty"`
	Status             string  `json:"status,omitempty"`
	Notes              string  `json:"notes,omitempty"`
	PromoID            int64   `json:"promoid,omitempty"`
	AutoRecalc         bool    `json:"autorecalc,omitempty"`
	UpdateNS           bool    `json:"updatens,omitempty"`
	NS1                string  `json:"ns1,omitempty"`
	NS2                string  `json:"ns2,omitempty"`
	NS3                string  `json:"ns3,omitempty"`
	NS4                string  `json:"ns4,omitempty"`
	NS5                string  `json:"ns5,omitempty"`
}

// SetRegDate sets the RegDate in the format expected by WHMCS using the given time t
func (c *ClientDomainUpdateReq) SetRegDate(t time.Time) {
	c.RegDate = t.Format("2006-01-02")
}

// SetNextDueDate sets the DueDate in the format expected by WHMCS using the given time t
func (c *ClientDomainUpdateReq) SetNextDueDate(t time.Time) {
	c.NextDueDate = t.Format("2006-01-02")
}

// SetExpiryDate sets the ExpiryDate in the format expected by WHMCS using the given time t
func (c *ClientDomainUpdateReq) SetExpiryDate(t time.Time) {
	c.ExpiryDate = t.Format("2006-01-02")
}

type ClientDomainUpdateResp struct {
	Result   string `json:"result"`
	DomainID int64  `json:"domainid,string"`
}

// UpdateClientDomain sends the updateclientdomain command with the given params.
func (a *API) UpdateClientDomain(d *ClientDomainUpdateReq) (r *ClientDomainUpdateResp, err error) {
	body, err := a.Do("updateclientdomain", &d)
	if err != nil {
		err = fmt.Errorf("gowhmcs updateclientdomain error: %v", err)
		return
	}

	r = &ClientDomainUpdateResp{}
	if err = json.Unmarshal(body, r); err != nil {
		err = fmt.Errorf("gowhmcs updateclientdomain error: %v", err)
		return
	}

	return
}
