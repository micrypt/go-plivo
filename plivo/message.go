// Public Domain (-) 2013-2014 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

type MessageService struct {
	client *Client
}

type MessageSendParams struct {
	Src  string `json:"src,omitempty"`
	Dst  string `json:"dst,omitempty"`
	Text string `json:"text,omitempty"`
	// Optional parameters.
	Type   string `json:"type,omitempty"`
	URL    string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
}

type Message struct {
	ToNumber         string `json:"to_number,omitempty"`
	FromNumber       string `json:"from_number,omitempty"`
	CloudRate        string `json:"cloud_rate,omitempty"`
	MessageType      string `json:"message_type,omitempty"`
	ResourceURI      string `json:"resource_uri,omitempty"`
	CarrierRate      string `json:"carrier_rate,omitempty"`
	MessageDirection string `json:"message_direction,omitempty"`
	MessageState     string `json:"message_state,omitempty"`
	TotalAmount      string `json:"total_amount,omitempty"`
	MessageUUID      string `json:"message_uuid,omitempty"`
	MessageTime      string `json:"message_time,omitempty"`
}

// Stores response for ending a message.
type MessageSendResponseBody struct {
	Message     string `json:"message"`
	ApiID       string `json:"api_id"`
	MessageUUID string `json:"message_uuid"`
}

// Make creates a call.
func (c *MessageService) Send(mp *MessageSendParams) (*Response, error) {
	req, err := c.client.NewRequest("POST", c.client.authID+"/Message/", mp)
	if err != nil {
		return nil, err
	}
	aResp := &MessageSendResponseBody{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req, aResp)
	return resp, err
}

type MessageGetAllParams struct {
	Limit  int64 `json:"limit:omitempty"`
	Offset int64 `json:"offset:omitempty"`
}

type MessageGetAllResponseBody struct {
	ApiID   string     `json:"api_id"`
	Meta    *Meta      `json:"meta"`
	Objects []*Message `json:"objects"`
}

// GetAll fetches all messages.
func (s *MessageService) GetAll(p *MessageGetAllParams) ([]*Message, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Message/", p)
	if err != nil {
		return nil, nil, err
	}
	aResp := &MessageGetAllResponseBody{}
	resp, err := s.client.Do(req, aResp)
	resp.Meta = aResp.Meta
	return aResp.Objects, resp, err
}

// Get fetches a specified message.
func (s *MessageService) Get(id string) (*Message, *Response, error) {
	req, err := s.client.NewRequest("GET", s.client.authID+"/Message/"+id+"/", nil)
	if err != nil {
		return nil, nil, err
	}
	aResp := &Message{}
	resp, err := s.client.Do(req, aResp)
	return aResp, resp, err
}
