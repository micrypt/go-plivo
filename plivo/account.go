// Public Domain (-) 2013 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type AccountService struct {
	client *Client
}

type Plan struct {
	VoiceRate           string `json:"voice_rate,omitempty"`
	MessagingRate       string `json:"messaging_rate,omitempty"`
	Name                string `json:"name_rate,omitempty"`
	MonthlyCloudCredits string `json:"voice_rate,omitempty"`
}

type Account struct {
	Address         string `json:"address,omitempty"`
	ApiID           string `json:"api_id,omitempty"`
	AuthID          string `json:"auth_id,omitempty"`
	AutoRecharge    bool   `json:"auto_recharge,omitempty"`
	BillingMode     string `json:"billing_mode,omitempty"`
	CashCredits     string `json:"cash_credits,omitempty"`
	CloudCredits    string `json:"cloud_credits,omitempty"`
	City            string `json:"city,omitempty"`
	CpsAllowed      string `json:"cps_allowed,omitempty"`
	Created         string `json:"created,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	GwType          string `json:"gw_type,omitempty"`
	Modified        string `json:"modified,omitempty"`
	Name            string `json:"name,omitempty"`
	Plan            Plan   `json:"plan,omitempty"`
	RechargeChoices string `json:"recharge_choices,omitempty"`
	ResourceURI     string `json:"resource_uri,omitempty"`
	State           string `json:"state,omitempty"`
	Timezone        string `json:"timezone,omitempty"`
}

type Subaccount struct {
	Account     string `json:"account,omitempty"`
	ApiID       string `json:"api_id,omitempty"`
	AuthID      string `json:"auth_id,omitempty"`
	AuthToken   string `json:"auth_token,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Name        string `json:"name,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// Get fetches an account.
func (s *AccountService) Get() (*Account, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/", nil)
	if err != nil {
		return nil, nil, err
	}

	aResp := new(Account)
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Stores response for Modify call
type ModifyResponseBody struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
}

// Modify edits an account
func (s *AccountService) Modify(acc *Account) (*Account, *Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/", acc)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	aResp := &ModifyResponseBody{}
	resp, err := s.client.Do(req, aResp)

	return acc, resp, err
}

// Stores response for CreateSubaccount call
type CreateResponseBody struct {
	AuthToken string `json:"auth_token"`
	Message   string `json:"message"`
	ApiID     string `json:"api_id"`
	AuthID    string `json:"auth_id"`
}

// CreateSubaccount creates a subaccount
func (s *AccountService) CreateSubaccount(sacc *Subaccount) (*Response, error) {

	req, err := s.client.NewRequest("POST", s.client.authID+"/Subaccount/", sacc)
	if err != nil {
		return nil, err
	}

	aResp := &CreateResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	sacc.AuthID = aResp.AuthID
	return resp, err
}

// ModifySubaccount edits a subaccount
func (s *AccountService) ModifySubaccount(sacc *Subaccount) (*Subaccount, *Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/Subaccount/"+sacc.AuthID+"/", sacc)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	aResp := &ModifyResponseBody{}
	resp, err := s.client.Do(req, aResp)

	return sacc, resp, err
}

// GetSubaccount fetches a subaccount.
func (s *AccountService) GetSubaccount(subAuthID string) (*Subaccount, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Subaccount/"+subAuthID+"/", nil)
	if err != nil {
		return nil, nil, err
	}

	aResp := &Subaccount{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

type SubaccountsResponseBody struct {
	ApiID   string        `json:"api_id"`
	Meta    *Meta         `json:"meta"`
	Objects []*Subaccount `json:"objects"`
}

// GetSubaccount fetches all subaccounts.
func (s *AccountService) GetSubaccounts(limit, offset int64) ([]*Subaccount, *Response, error) {
	limitOffset := &limitOffset{limit, offset}

	req, err := s.client.NewRequest("GET", s.client.authID+"/Subaccount/", limitOffset)

	if err != nil {
		return nil, nil, err
	}

	aResp := &SubaccountsResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// DeleteSubaccount deletes a subaccount.
func (s *AccountService) DeleteSubaccount(subAuthID string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Subaccount/"+subAuthID+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}
