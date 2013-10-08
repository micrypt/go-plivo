package plivo

type EndpointService struct {
	client *Client
}

type Endpoint struct {
	Alias       string `json:"alias,omitempty"`
	EndpointID  string `json:"endpoint_id,omitempty"`
	Password    string `json:"password,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	SIPURI      string `json:"sip_uri,omitempty"`
	Username    string `json:"username,omitempty"`
	// Optional field for Create call.
	AppID string `json:"app_id,omitempty"`
}

type EndpointsResponseBody struct {
	ApiID   string      `json:"api_id"`
	Meta    *Meta       `json:"meta"`
	Objects []*Endpoint `json:"objects"`
}

// GetEndpoints retrieves a list of all endpoints.
func (s *EndpointService) GetEndpoints(limit, offset int64) ([]*Endpoint, *Response, error) {
	limitOffset := &limitOffset{limit, offset}

	req, err := s.client.NewRequest("GET", s.client.authID+"/Endpoint/", limitOffset)

	if err != nil {
		return nil, nil, err
	}

	aResp := &EndpointsResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Stores response for Create call
type EndpointCreateResponseBody struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
}

// Create creates an endpoint.
func (s *EndpointService) Create(ep *Endpoint) (*Endpoint, *Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/Endpoint/", ep)
	if err != nil {
		return nil, nil, err
	}
	aResp := &EndpointCreateResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	return ep, resp, err
}

// Get fetches a particular endpoint.
func (s *EndpointService) Get(id string) (*Endpoint, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Endpoint/"+id+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Endpoint{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Modify edits an endpoint.
func (s *EndpointService) Modify(ep *Endpoint) (*Endpoint, *Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/Endpoint/"+ep.EndpointID+"/", ep)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	aResp := &ModifyResponseBody{}
	resp, err := s.client.Do(req, aResp)

	return ep, resp, err
}

// Delete deletes an endpoint.
func (s *EndpointService) Delete(id string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Endpoint/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}
