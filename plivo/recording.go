// Public Domain (-) 2013-2014 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type RecordingService struct {
	client *Client
}

type Recording struct {
	CallUUID            string `json:"call_uuid,omitempty"`
	RecordingID         string `json:"recording_id,omitempty"`
	RecordingType       string `json:"recording_type,omitempty"`
	RecordingFormat     string `json:"recording_format,omitempty"`
	ConferenceName      string `json:"conference_name,omitempty"`
	RecordingURL        string `json:"recording_url,omitempty"`
	ResourceURI         string `json:"resource_uri,omitempty"`
	RecordingStartMS    string `json:"recording_start_ms,omitempty"`
	RecordingEndMS      string `json:"recording_end_ms,omitempty"`
	RecordingDurationMS string `json:"recording_duration_ms,omitempty"`
}

type RecordingGetAllParams struct {
	// Query parameters.
	Subaccount string `json:"subaccount,omitempty"`
	CallUUID   string `json:"call_uuid,omitempty"`
	AddTime    string `json:"add_time,omitempty"`
	Limit      int64  `json:"limit:omitempty"`
	Offset     int64  `json:"offset:omitempty"`
}

type RecordingGetAllResponseBody struct {
	ApiID   string       `json:"api_id"`
	Meta    *Meta        `json:"meta"`
	Objects []*Recording `json:"objects"`
}

// GetAll fetches all recordings.
func (s *RecordingService) GetAll(p *RecordingGetAllParams) ([]*Recording, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Recording/", p)
	if err != nil {
		return nil, nil, err
	}
	aResp := &RecordingGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified recording.
func (s *RecordingService) Get(recordingID string) (*Recording, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Recording/"+recordingID+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Recording{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}
