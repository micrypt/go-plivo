package plivo

type ApplicationService struct {
	client *Client
}

type Application struct {
	FallbackMethod    string `json:"fallback_method,omitempty"`
	DefaultApp        bool   `json:"default_app,omitempty"`
	AppName           string `json:"app_name,omitempty"`
	ProductionApp     bool   `json:"production_app,omitempty"`
	AppID             string `json:"app_id,omitempty"`
	HangupURL         string `json:"hangup_url,omitempty"`
	AnswerURL         string `json:"answer_url,omitempty"`
	MessageURL        string `json:"message_url,omitempty"`
	ResourceURI       string `json:"resource_uri,omitempty"`
	HangupMethod      string `json:"hangup_method,omitempty"`
	MessageMethod     string `json:"message_method,omitempty"`
	FallbackAnswerURL string `json:"fallback_answer_url,omitempty"`
	AnswerMethod      string `json:"answer_method,omitempty"`
	ApiID             string `json:"api_id,omitempty"`

	// Additional fields for Modify calls
	DefaultNumberApp   bool `json:"default_number_app,omitempty"`
	DefaultEndpointApp bool `json:"default_endpoint_app,omitempty"`
}

// Stores response for Create call
type ApplicationCreateResponseBody struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
	AppID   string `json:"app_id"`
}

// CreateApplication creates an application.
func (s *ApplicationService) Create(app *Application) (*Application, *Response, error) {

	req, err := s.client.NewRequest("POST", s.client.authID+"/Application/", app)

	if err != nil {
		return nil, nil, err
	}

	aResp := &ApplicationCreateResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, aResp)
	app.AppID = aResp.AppID
	return app, resp, err
}

type ApplicationsResponseBody struct {
	ApiID   string         `json:"api_id"`
	Meta    *Meta          `json:"meta"`
	Objects []*Application `json:"objects"`
}

// GetApplication fetches all subaccounts.
func (s *ApplicationService) GetApplications(limit, offset int64) ([]*Application, *Response, error) {
	limitOffset := &limitOffset{limit, offset}

	req, err := s.client.NewRequest("GET", s.client.authID+"/Application/", limitOffset)

	if err != nil {
		return nil, nil, err
	}

	aResp := &ApplicationsResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified application.
func (s *ApplicationService) Get(appID string) (*Application, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Application/"+appID+"/", nil)
	if err != nil {
		return nil, nil, err
	}

	aResp := &Application{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Modify edits an application
func (s *ApplicationService) Modify(app *Application) (*Application, *Response, error) {
	req, err := s.client.NewRequest("POST", s.client.authID+"/Application/"+app.AppID+"/", app)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	aResp := &ModifyResponseBody{}
	resp, err := s.client.Do(req, aResp)

	return app, resp, err
}

// Delete deletes a subaccount.
func (s *ApplicationService) Delete(subAuthID string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Application/"+subAuthID+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}