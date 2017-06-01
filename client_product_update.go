package whmcs

import (
	"errors"
)

var (
	ErrNoProductID = errors.New("No pid specified")
)

type UpdateClientProductResult struct {
	Result    string `json:"result" xml:"result"`
	ServiceID int64  `json:"serviceid" xml:"serviceid"`
}

type ClientProduct struct {
	ServiceID            int64         `json:"serviceid" xml:"serviceid"`
	PID                  string        `json:"pid,omitempty" xml:"pid,omitempty"`
	ServerID             string        `json:"serverid,omitempty" xml:"serverid,omitempty"`
	RegDate              ProductDate   `json:"regdate,omitempty" xml:"regdate,omitempty"`
	NextDueDate          ProductDate   `json:"nextduedate,omitempty" xml:"nextduedate,omitempty"`
	Domain               string        `json:"domainomitempty,omitempty" xml:"domainomitempty,omitempty"`
	FirstPaymentAmount   float64       `json:"firstpaymentamount,string,omitempty" xml:"firstpaymentamount,omitempty"`
	RecurringAmount      float64       `json:"recurringamount,string,omitempty" xml:"recurringamount,omitempty"`
	BillingCycle         string        `json:"billingcycle,omitempty" xml:"billingcycle,omitempty"`
	PaymentMethod        string        `json:"paymentmethod,omitempty" xml:"paymentmethod,omitempty"`
	Status               string        `json:"status,omitempty" xml:"status,omitempty"`
	ServiceUsername      string        `json:"serviceusername,omitempty" xml:"serviceusername,omitempty"`
	ServicePassword      string        `json:"servicepassword,omitempty" xml:"servicepassword,omitempty"`
	SubscriptionID       string        `json:"subscriptionid,omitempty" xml:"subscriptionid,omitempty"`
	PromoID              string        `json:"promoid,omitempty" xml:"promoid,omitempty"`
	OverrideAutoSuspend  string        `json:"overideautosuspend,omitempty" xml:"overideautosuspend,omitempty"`
	OverrideSuspendUntil ProductDate   `json:"overidesuspenduntil,omitempty" xml:"overideautosuspend,omitempty"`
	NS1                  string        `json:"ns1,omitempty" xml:"ns1,omitempty"`
	NS2                  string        `json:"ns2,omitempty" xml:"ns2,omitempty"`
	DedicatedIP          string        `json:"dedicatedip,omitempty" xml:"dedicatedip,omitempty"`
	AssignedIPs          string        `json:"assignedips,omitempty" xml:"assignedips,omitempty"`
	Notes                string        `json:"notes,omitempty" xml:"notes,omitempty"`
	AutoRecalc           bool          `json:"autorecalc,string,omitempty" xml:"autorecalc,omitempty"`
	CustomFields         CustomFields  `json:"customfields,omitempty" xml:"customfields,omitempty"`
	ConfigOptions        ConfigOptions `json:"configoptions,omitempty" xml:"configoptions,omitempty"`
}

func (p *ClientProduct) Error() error {
	if p.ServiceID == 0 {
		return ErrNoProductID
	}
	return nil
}
