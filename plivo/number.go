// Public Domain (-) 2013 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type NumberService struct {
	client *Client
}

type Number struct {
	VoiceEnabled bool   `json:"voice_enabled,omitempty"`
	SMSEnabled   bool   `json:"sms_enabled,omitempty"`
	Description  string `json:"description,omitempty"`
	PlivoNumber  bool   `json:"plivo_number,omitempty"`
	Number       string `json:"number,omitempty"`
	NumberType   string `json:"number_type,omitempty"`
	Application  string `json:"application,omitempty"`
	AddedOn      string `json:"added_on,omitempty"`
	ResourceURI  string `json:"resource_uri,omitempty"`
}

type NumberGetAllParams struct {
	NumberType       string `json:"number_type,omitempty"`
	NumberStartswith string `json:"number_startswith,omitempty"`
	Subaccount       string `json:"subaccount,omitempty"`
	Services         string `json:"services,omitempty"`
	Limit            int64  `json:"limit:omitempty"`
	Offset           int64  `json:"offset:omitempty"`
}

type NumberGetAllResponseBody struct {
	ApiID   string    `json:"api_id"`
	Meta    *Meta     `json:"meta"`
	Objects []*Number `json:"objects"`
}

// GetAll fetches all calls.
func (s *NumberService) GetAll(p *NumberGetAllParams) ([]*Number, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Number/", p)

	if err != nil {
		return nil, nil, err
	}
	aResp := &NumberGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get gets details of a rented number.
func (s *NumberService) Get(number string) (*Number, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Number/"+number+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Number{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

type NumberAddParams struct {
	Numbers string `json:"numbers"`
	Carrier string `json:"carrier"`
	Region  string `json:"region"`
	// Optional parameters.
	NumberType string `json:"number_type,omitempty"`
	AppID      string `json:"app_id,omitempty"`
	Subaccount string `json:"subaccount,omitempty"`
}

// Add adds a number from your own carrier.
func (c *NumberService) Add(np *NumberAddParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Number/", np)
	if err != nil {
		return nil, err
	}
	aResp := &ModifyResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

type NumberEditParams struct {
	AppID      string `json:"app_id,omitempty"`
	Subaccount string `json:"subaccount,omitempty"`
}

// Add adds a number from your own carrier.
func (c *NumberService) Add(number string, np *NumberEditParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Number/"+number+"/", np)
	if err != nil {
		return nil, err
	}
	aResp := &ModifyResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// Unrent unrents a number.
func (s *NumberService) Unrent(number string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Number/"+number+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}
