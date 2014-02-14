// Public Domain (-) 2013-2014 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type OutgoingCarrierService struct {
	client *Client
}

type OutgoingCarrier struct {
	CarrierID   string
	IPSet       string
	Name        string
	ResourceURI string
}

type OutgoingCarrierGetAllResponseBody struct {
	ApiID   string             `json:"api_id"`
	Meta    *Meta              `json:"meta"`
	Objects []*OutgoingCarrier `json:"objects"`
}

type OutgoingCarrierGetAllParams struct {
	// Query parameters.
	Name   string `json:"name,omitempty"`
	Limit  int64  `json:"limit:omitempty"`
	Offset int64  `json:"offset:omitempty"`
}

func (s *OutgoingCarrierService) GetAll(p *OutgoingCarrierGetAllParams) ([]*OutgoingCarrier, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/OutgoingCarrier/", p)
	if err != nil {
		return nil, nil, err
	}
	aResp := &OutgoingCarrierGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified carrier.
func (s *OutgoingCarrierService) Get(carrierID string) (*OutgoingCarrier, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/OutgoingCarrier/"+carrierID+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &OutgoingCarrier{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Remove removes a carrier, and deletes all numbers associated with the carrier.
func (s *OutgoingCarrierService) Remove(carrierID string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/OutgoingCarrier/"+carrierID+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}

type OutgoingCarrierAddParams struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	FailoverAddress string `json:"failover_address,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	FailoverPrefix  string `json:"failover_prefix,omitempty"`
	Suffix          string `json:"suffix,omitempty"`
	Retries         int64  `json:"retries,omitempty"`
	RetrySeconds    int64  `json:"retry_seconds,omitempty"`
}

// Stores response for making a call.
type OutgoingCarrierResponseBody struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
}

// Add adds an outgoing carrier.
func (s *OutgoingCarrierService) Add(p *OutgoingCarrierAddParams) (*Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/OutgoingCarrier/", p)
	if err != nil {
		return nil, err
	}
	aResp := &OutgoingCarrierResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	return resp, err
}

type OutgoingCarrierModifyParams struct {
	Name  string `json:"name,omitempty"`
	IPSet string `json:"ip_set,omitempty"`
}

// Modify updates an outgoing carrier.
func (s *OutgoingCarrierService) Modify(p *OutgoingCarrierModifyParams) (*Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/OutgoingCarrier/", p)
	if err != nil {
		return nil, err
	}
	aResp := &OutgoingCarrierResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	return resp, err
}
