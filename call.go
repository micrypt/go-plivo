package plivo

type CallService struct {
	client *Client
}

type CallMakeParams struct {
	// Required parameters.
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	AnswerUrl string `json:"answer_url,omitempty"`
	// Optional parameters.
	AnswerMethod         string `json:"answer_method,omitempty"`
	RingUrl              string `json:"ring_url,omitempty"`
	RingMethod           string `json:"ring_method,omitempty"`
	HangupUrl            string `json:"hangup_url,omitempty"`
	HangupMethod         string `json:"hangup_method,omitempty"`
	FallbackUrl          string `json:"fallback_url,omitempty"`
	FallbackMethod       string `json:"fallback_method,omitempty"`
	CallerName           string `json:"caller_name,omitempty"`
	SendDigits           string `json:"send_digits,omitempty"`
	SendOnPreanswer      bool   `json:"send_on_preanswer,omitempty"`
	TimeLimit            int64  `json:"time_limit,omitempty"`
	HangupOnRing         int64  `json:"hangup_on_ring,omitempty"`
	MachineDetection     string `json:"machine_detection,omitempty"`
	MachineDetectionTime int64  `json:"machine_detection_time,omitempty"`
	SipHeaders           string `json:"sip_headers,omitempty"`
	RingTimeout          int64  `json:"ring_timeout,omitempty"`
}

type Call struct {
	FromNumber     string `json:"from_number,omitempty"`
	ToNumber       string `json:"to_number,omitempty"`
	AnswerUrl      string `json:"answer_url,omitempty"`
	CallUUID       string `json:"call_uuid,omitempty"`
	ParentCallUUID string `json:"parent_call_uuid,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	TotalAmount    string `json:"total_amount,omitempty"`
	CallDirection  string `json:"call_direction,omitempty"`
	CallDuration   int64  `json:"call_duration,omitempty"`
	MessageUrl     string `json:"message_url,omitempty"`
	ResourceUri    string `json:"resource_uri,omitempty"`
}

type LiveCall struct {
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
	AnswerUrl      string `json:"answer_url,omitempty"`
	CallUUID       string `json:"call_uuid,omitempty"`
	CallerName     string `json:"caller_name,omitempty"`
	ParentCallUUID string `json:"parent_call_uuid,omitempty"`
	SessionStart   string `json:"session_start,omitempty"`
}

// Stores response for making a call.
type CallMakeResponseBody struct {
	Message  string `json:"message"`
	ApiID    string `json:"api_id"`
	AppID    string `json:"app_id"`
	CallUUID string `json:"call_uuid"`
}

// Make creates a call.
func (c *CallService) Make(cp *CallMakeParams) (*Response, error) {

	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/", cp)

	if err != nil {
		return nil, err
	}

	aResp := &CallMakeResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

type CallGetAllParams struct {
	// Query parameters.
	SubAccount    string `json:"subaccount,omitempty"`
	CallDirection string `json:"call_direction,omitempty"`
	FromNumber    string `json:"from_number,omitempty"`
	ToNumber      string `json:"to_number,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	BillDuration  string `json:"bill_duration,omitempty"`
	Limit         int64  `json:"limit:omitempty"`
	Offset        int64  `json:"offset:omitempty"`
}

type CallGetAllResponseBody struct {
	ApiID   string  `json:"api_id"`
	Meta    *Meta   `json:"meta"`
	Objects []*Call `json:"objects"`
}

// GetAll fetches all calls.
func (s *CallService) GetAll(p *CallGetAllParams) ([]*Call, *Response, error) {

	req, err := s.client.NewRequest("GET", s.client.authID+"/Call/", p)

	if err != nil {
		return nil, nil, err
	}

	aResp := &CallGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified call.
func (s *CallService) Get(callID string) (*Call, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Call/"+callID+"/", nil)
	if err != nil {
		return nil, nil, err
	}

	aResp := &Call{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// GetCallLive fetches all live calls.
func (s *CallService) GetAllLive() ([]*Call, *Response, error) {

	req, err := s.client.NewRequest("GET", s.client.authID+"/Call/?status=live", nil)

	if err != nil {
		return nil, nil, err
	}

	aResp := &CallGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// GetLive fetches details of a specified call.
func (s *CallService) GetLive(uuid string) (*LiveCall, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Call/"+uuid+"/?status=live", nil)
	if err != nil {
		return nil, nil, err
	}

	aResp := &LiveCall{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// Hangup terminates a specified call.
func (s *CallService) Hangup(uuid string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Call/"+uuid+"/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}

type CallTransferParams struct {
	legs       string `json:"legs,omitempty"`
	AlegUrl    string `json:"aleg_url,omitempty"`
	AlegMethod string `json:"aleg_method,omitempty"`
	BlegUrl    string `json:"bleg_url,omitempty"`
	BlegMethod string `json:"bleg_method,omitempty"`
}

type CallTransferResponseBody struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message"`
}

// Transfer transfers a call.
func (c *CallService) Transfer(cp *CallTransferParams) (*Response, error) {

	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/", cp)

	if err != nil {
		return nil, err
	}

	aResp := &CallTransferResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

type CallRecordParams struct {
	TimeLimit           int64  `json:"time_limit,omitempty"`
	FileFormat          string `json:"file_format,omitempty"`
	TranscriptionType   string `json:"transcription_type,omitempty"`
	TranscriptionUrl    string `json:"transcription_url,omitempty"`
	TranscriptionMethod string `json:"transcription_method,omitempty"`
	CallbackUrl         string `json:"callback_url,omitempty"`
	CallbackMethod      string `json:"callback_method,omitempty"`
}

type CallRecordResponseBody struct {
	Message string `json:"message,omitempty"`
	Url     string `json:"url,omitempty"`
}

// Record records a call.
func (c *CallService) Record(uuid string, cp *CallRecordParams) (*Response, error) {

	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/"+uuid+"/Record/", cp)

	if err != nil {
		return nil, err
	}

	aResp := &CallRecordResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// Stop cancels a call recording.
func (c *CallService) StopRecording(uuid, url string) (*Response, error) {

	rp := struct{ URL string }{url}

	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/"+uuid+"/Record/", rp)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, nil)
	return resp, err
}