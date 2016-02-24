package whmcs

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/bradberger/gowhmcs/serializer"
)

var (
	ErrNoAPIURL      = errors.New("THe WHMCS API URL endpoint is empty")
	ErrNoAPIUsername = errors.New("THe WHMCS API username is empty")
)

// CustomFields @TODO
type CustomFields map[string]string

// MarshalJSON formats the custom fields in a way that WHMCS recognizes them.
func (c CustomFields) MarshalJSON() ([]byte, error) {
	str, err := serializer.Encode(c)
	if err != nil {
		return []byte{}, err
	}
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return []byte(enc), nil
}

// ConfigOptions @TODO
type ConfigOptions map[string]string

// MarshalJSON formats the custom fields in a way that WHMCS recognizes them.
func (c ConfigOptions) MarshalJSON() ([]byte, error) {
	str, err := serializer.Encode(c)
	if err != nil {
		return []byte{}, err
	}
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return []byte(enc), nil
}

func NewAPI(url, user, pwd string) (api *API, err error) {

	if url == "" {
		err = ErrNoAPIURL
		return
	}

	if user == "" {
		err = ErrNoAPIURL
		return
	}

	if pwd == "" {
		err = ErrNoAPIURL
		return
	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	api = &API{
		Endpoint: url,
		Username: user,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(pwd))),
	}

	return

}

type APIResponse struct {
	Message string `json:"message" xml:"message"`
	Result  string `json:"result" xml:"result"`
}

type APIBasicResponse struct {
	Result string `json:"result" xml:"result"`
}

func (r *APIBasicResponse) Success() bool {
	return r.Result == "success"
}

func (r *APIResponse) Error() error {
	if r.Result == "error" {
		return errors.New(r.Message)
	}
	return nil
}

func (a *API) Do(cmd string, data interface{}) ([]byte, error) {

	// Encode the data to send.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	// Now encode that into a map[string]string to change to url.Values{}
	m := make(map[string]string)
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		return []byte{}, err
	}

	m["responsetype"] = "json"
	m["username"] = a.Username
	m["password"] = a.Password
	m["action"] = cmd

	form := url.Values{}
	for k, v := range m {
		form.Add(k, v)
	}

	// POST it.
	log.Printf("Send: %+v", m)
	url := fmt.Sprintf("%s/includes/api.php", a.Endpoint)
	r, err := http.PostForm(url, form)
	if err != nil {
		return []byte{}, err
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}

	log.Printf("Body: %s", body)

	// The most basic responses have no message, so allow for that here.
	s := APIBasicResponse{}
	if err := json.Unmarshal(body, &s); err != nil {
		return body, err
	}

	// We do this to first check for errors.
	// If not success, it will have a message, so parse that and return it here.
	if !s.Success() {

		e := APIResponse{}
		if err := json.Unmarshal(body, &e); err != nil {
			return body, err
		}

		if err := e.Error(); err != nil {
			return body, err
		}

	}

	return body, nil

}
