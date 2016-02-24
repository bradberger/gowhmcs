package whmcs

type WhoisInfo struct {
	Result string `json:"result" xml:"result"`
	Status string `json:"status" xml:"status"`
	Whois  string `json:"whois" xml:"whois"`
}

type WhoisRequest struct {
	Action string `json:"action" xml:"action"`
	Domain string `json:"domain" xml:"domain"`
}
