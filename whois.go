package whmcs

import "encoding/json"

type WhoisInfo struct {
	Result string `json:"result" xml:"result"`
	Status string `json:"status" xml:"status"`
	Whois  string `json:"whois" xml:"whois"`
}

type WhoisRequest struct {
	Action string `json:"action" xml:"action"`
	Domain string `json:"domain" xml:"domain"`
}

func (a *API) Whois(domain string) (w WhoisInfo, err error) {

	body, err := a.Do("domainwhois", &WhoisRequest{"domainwhois", domain})
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &w)
	return

}
