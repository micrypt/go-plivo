package plivo

type ConferenceService struct {
	client *Client
}

type Conference struct {
	ConferenceName        string   `json:"conference_name,omitempty"`
	ConferenceRunTime     string   `json:"conference_run_time,omitempty"`
	ConferenceMemberCount string   `json:"conference_member_count,omitempty"`
	Members               []Member `json:"members,omitempty"`
}

type Member struct {
	Muted      bool   `json:"muted,omitempty"`
	MemberID   string `json:"member_id,omitempty"`
	Deaf       bool   `json:"deaf,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	CallerName string `json:"caller_name,omitempty"`
	Direction  string `json:"direction,omitempty"`
	CallUUID   string `json:"call_uuid,omitempty"`
	JoinTime   string `json:"join_time,omitempty"`
}

type ConferenceGetAllAllResponseBody struct {
	ApiID       string   `json:"api_id"`
	Conferences []string `json:"conferences"`
}

// GetAll retrieves list of all conferences.
func (s *ConferenceService) GetAll() ([]string, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Conference/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &ConferenceGetAllAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	return aResp.Conferences, resp, err
}

// Get retrieves details of a particular conference.
func (s *ConferenceService) Get(name string) (*Conference, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Conference/"+name+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Conference{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}

// HangupAll hangs up all conferences.
func (s *ConferenceService) HangupAll() (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}

// Hangup hangs up a particular conference.
func (s *ConferenceService) Hangup(name string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}

// HangupMember hangs up member(s).
func (s *ConferenceService) HangupMember(name, member string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/Member/"+member+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

// KickMembers kicks member(s).
func (s *ConferenceService) KickMembers(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Kick/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

// MuteMembers mutes member(s).
func (s *ConferenceService) MuteMembers(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("POST", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Mute/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

// UnmuteMembers unmutes member(s).
	func (s *ConferenceService) UnmuteMembers(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Mute/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

// Play starts playing sound to member(s).
func (s *ConferenceService) Play(name, members, url string) (*Response, error) {	
	rp := struct{ URL string }{url}
	req, err := s.client.NewRequest("POST", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Play/", rp)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

// StopPlaying stops playing sound to member(s).
func (s *ConferenceService) StopPlaying(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Play/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

type ConferenceSpeakParams struct {
	Text     string `json:"text"`
	Voice    string `json:"length,omitempty"`
	Language string `json:"language,omitempty"`
}

type ConferenceSpeakResponseBody struct {
	Message string `json:"message,omitempty"`
	ApiID   string `json:"api_id,omitempty"`
}

// Speak makes member(s) listen to a speech.
func (c *ConferenceService) Speak(name, members string, cp *ConferenceSpeakParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Conference/"+name+"/Member/"+members+"/Speak/", cp)
	if err != nil {
		return nil, err
	}
	aResp := &ConferenceSpeakResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// DisableHearingMembers makes member(s) deaf.
func (s *ConferenceService) DisableHearingMembers(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("POST", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Deaf/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

//EnableHearingMembers enables hearing for member(s).
func (s *ConferenceService) EnableHearingMembers(name, members string) (*Response, error) {	
	req, err := s.client.NewRequest("DELETE", s.client.authID+"/Conference/"+name+"/Member/"+members+"/Deaf/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err	
}

type ConferenceRecordParams struct {
	TimeLimit           int64  `json:"time_limit,omitempty"`
	FileFormat          string `json:"file_format,omitempty"`
	TranscriptionType   string `json:"transcription_type,omitempty"`
	TranscriptionUrl    string `json:"transcription_url,omitempty"`
	TranscriptionMethod string `json:"transcription_method,omitempty"`
	CallbackUrl         string `json:"callback_url,omitempty"`
	CallbackMethod      string `json:"callback_method,omitempty"`
}

type ConferenceRecordResponseBody struct {
	Message string `json:"message,omitempty"`
	Url     string `json:"url,omitempty"`
}

// Record records a conference.
func (c *ConferenceService) Record(id string, cp *ConferenceRecordParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Conference/"+id+"/Record/", cp)
	if err != nil {
		return nil, err
	}
	aResp := &ConferenceRecordResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

// StopRecording cancels a conference recording.
func (c *ConferenceService) StopRecording(id string) (*Response, error) {
	req, err := c.client.NewRequest("DELETE", c.client.authID+"/Conference/"+id+"/Record/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req, nil)
	return resp, err
}