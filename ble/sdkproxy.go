package ble

import (
	"encoding/json"
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// SDKProxyRequest is the type required for SDK proxy messages
type SDKProxyRequest struct {
	URLPath string
	Body    string
}

// SDKProxyResponse is the response type of an SDKProxy request
type SDKProxyResponse struct {
	MessageID    string `json:"message_id,omitempty"`
	StatusCode   uint16 `json:"status_code,omitempty"`
	ResponseType string `json:"response_type,omitempty"`
	ResponseBody string `json:"response_body,omitempty"`
}

// Marshal converts a SDKProxyResponse message to bytes
func (sr *SDKProxyResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a SDKProxyResponse byte slice to a SDKProxyResponse
func (sr *SDKProxyResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// SDKProxy sends a BLE-tunneled SDK request
func (v *VectorBLE) SDKProxy(settings *SDKProxyRequest) (*SDKProxyResponse, error) {
	if !v.state.authorized || v.state.clientGUID == "" {
		return nil, errors.New(errNotAuthorized)
	}

	if v.ble.Version() != rtsV5 {
		return nil, errors.New("unsupported version")
	}

	msg, _ := rts5.BuildSDKMessage(
		v.state.clientGUID,
		"1",
		settings.URLPath,
		settings.Body,
	)

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	b, err := v.watch()

	resp := SDKProxyResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err
}

func handleRST5SDKProxyResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	m := msg.GetRtsSdkProxyResponse()

	resp := SDKProxyResponse{
		MessageID:    m.MessageId,
		StatusCode:   m.StatusCode,
		ResponseType: m.ResponseType,
		ResponseBody: m.ResponseBody,
	}

	b, err := resp.Marshal()
	return b, false, err
}
