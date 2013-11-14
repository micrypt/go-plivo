// Public Domain (-) 2013 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type CallService struct {
	client *Client
}

type CallMakeParams struct {
	// Required parameters.
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	AnswerURL string `json:"answer_url,omitempty"`
	// Optional parameters.
	AnswerMethod         string `json:"answer_method,omitempty"`
	RingURL              string `json:"ring_url,omitempty"`
	RingMethod           string `json:"ring_method,omitempty"`
	HangupURL            string `json:"hangup_url,omitempty"`
	HangupMethod         string `json:"hangup_method,omitempty"`
	FallbackURL          string `json:"fallback_url,omitempty"`
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
	AnswerURL      string `json:"answer_url,omitempty"`
	CallUUID       string `json:"call_uuid,omitempty"`
	ParentCallUUID string `json:"parent_call_uuid,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	TotalAmount    string `json:"total_amount,omitempty"`
	CallDirection  string `json:"call_direction,omitempty"`
	CallDuration   int64  `json:"call_duration,omitempty"`
	MessageURL     string `json:"message_url,omitempty"`
	ResourceURI    string `json:"resource_uri,omitempty"`
}

type LiveCall struct {
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
	AnswerURL      string `json:"answer_url,omitempty"`
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
	Subaccount    string `json:"subaccount,omitempty"`
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
	AlegURL    string `json:"aleg_url,omitempty"`
	AlegMethod string `json:"aleg_method,omitempty"`
	BlegURL    string `json:"bleg_url,omitempty"`
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
	TranscriptionURL    string `json:"transcription_url,omitempty"`
	TranscriptionMethod string `json:"transcription_method,omitempty"`
	CallbackURL         string `json:"callback_url,omitempty"`
	CallbackMethod      string `json:"callback_method,omitempty"`
}

type CallRecordResponseBody struct {
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
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

// StopRecording cancels a call recording.
func (c *CallService) StopRecording(uuid, url string) (*Response, error) {
	rp := struct{ URL string }{url}
	req, err := c.client.NewRequest("DELETE", c.client.authID+"/Call/"+uuid+"/Record/", rp)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, nil)
	return resp, err
}

type CallPlayParams struct {
	URLs   string `json:"urls"`
	Length string `json:"length,omitempty"`
	Legs   string `json:"legs,omitempty"`
	Loop   bool   `json:"loop,omitempty"`
	Mix    bool   `json:"mix,omitempty"`
}

type CallPlayResponseBody struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// Play plays and controls sounds during a call.
func (c *CallService) Play(uuid string, cp *CallPlayParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/"+uuid+"/Play/", cp)
	if err != nil {
		return nil, err
	}
	aResp := &CallPlayResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// StopPlaying stops playing sounds during a call.
func (c *CallService) StopPlaying(uuid string) (*Response, error) {
	req, err := c.client.NewRequest("DELETE", c.client.authID+"/Call/"+uuid+"/Play/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, nil)
	return resp, err
}

type CallSpeakParams struct {
	Text     string `json:"text"`
	Voice    string `json:"length,omitempty"`
	Language string `json:"language,omitempty"`
	Legs     string `json:"legs,omitempty"`
	Loop     bool   `json:"loop,omitempty"`
	Mix      bool   `json:"mix,omitempty"`
}

type CallSpeakResponseBody struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// Speak plays text during a call (text to speech).
func (c *CallService) Speak(uuid string, cp *CallSpeakParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/"+uuid+"/Speak/", cp)
	if err != nil {
		return nil, err
	}
	aResp := &CallSpeakResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// StopSpeaking stops playing text during a call.
func (c *CallService) StopSpeaking(uuid string) (*Response, error) {
	req, err := c.client.NewRequest("DELETE", c.client.authID+"/Call/"+uuid+"/Speak/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, nil)
	return resp, err
}

type CallDTMFParams struct {
	Digits string `json:"digits"`
	Legs   string `json:"legs,omitempty"`
}

type CallDTMFResponseBody struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// DTMF send digits on a call.
func (c *CallService) DTMF(uuid string, cp *CallDTMFParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Call/"+uuid+"/DTMF/", cp)
	if err != nil {
		return nil, err
	}
	aResp := &CallDTMFResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// Cancel hangups a call request.
func (c *CallService) Cancel(request_uuid string) (*Response, error) {
	req, err := c.client.NewRequest("DELETE", c.client.authID+"/Request/"+request_uuid, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, nil)
	return resp, err
}
