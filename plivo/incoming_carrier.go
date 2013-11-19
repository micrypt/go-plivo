// Public Domain (-) 2013 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type IncomingCarrierService struct {
	client *Client
}

type IncomingCarrier struct {
	CarrierID   string
	IPSet       string
	Name        string
	ResourceURI string
	SMS         bool
	Voice       bool
}

type IncomingCarrierGetAllResponseBody struct {
	ApiID   string             `json:"api_id"`
	Meta    *Meta              `json:"meta"`
	Objects []*IncomingCarrier `json:"objects"`
}

type IncomingCarrierGetAllParams struct {
	// Query parameters.
	Name   string `json:"name,omitempty"`
	Limit  int64  `json:"limit:omitempty"`
	Offset int64  `json:"offset:omitempty"`
}

func (s *IncomingCarrierService) GetAll(p *IncomingCarrierGetAllParams) ([]*IncomingCarrier, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/IncomingCarrier/", p)
	if err != nil {
		return nil, nil, err
	}
	aResp := &IncomingCarrierGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified carrier.
func (s *IncomingCarrierService) Get(carrierID string) (*IncomingCarrier, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/IncomingCarrier/"+carrierID+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &IncomingCarrier{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Remove removes a carrier, and deletes all numbers associated with the carrier.
func (s *CallService) Remove(carrierID string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/IncomingCarrier/"+carrierID+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}

type IncomingCarrierAddParams struct {
	Name  string `json:"name"`
	IPSet string `json:"ip_set"`
}

// Stores response for making a call.
type IncomingCarrierResponseBody struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
}

// Add adds an incoming carrier.
func (s *IncomingCarrierService) Add(p *IncomingCarrierAddParams) (*Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/IncomingCarrier/", p)
	if err != nil {
		return nil, err
	}
	aResp := &IncomingCarrierResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	return resp, err
}

type IncomingCarrierModifyParams struct {
	Name  string `json:"name,omitempty"`
	IPSet string `json:"ip_set,omitempty"`
}

// Modify updates an incoming carrier.
func (s *IncomingCarrierService) Modify(p *IncomingCarrierModifyParams) (*Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/IncomingCarrier/", p)
	if err != nil {
		return nil, err
	}
	aResp := &IncomingCarrierResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	return resp, err
}
