package ble

import (
	"encoding/json"
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// AuthStatus defines the status of the response message
type AuthStatus int

const (
	UnknownError AuthStatus = iota
	ConnectionError
	WrongAccount
	InvalidSessionToken
	AuthorizedAsPrimary
	AuthorizedAsSecondary
	Reauthorized
)

// AuthResponse is the unified response for Auth  messages
type AuthResponse struct {
	Status          AuthStatus `json:"status,omitempty"`
	ClientTokenGUID string     `json:"client_token_guid,omitempty"`
	Success         bool       `json:"success,omitempty"`
}

// Marshal converts a AuthResponse message to bytes
func (sr *AuthResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a AuthResponse byte slice to a AuthResponse
func (sr *AuthResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// Auth sends a Auth message to the vector robot
func (v *VectorBLE) Auth(key string) (*AuthResponse, error) {
	var (
		msg []byte
		err error
	)

	switch v.ble.Version() {
	case rtsV3:
		msg, err = rts3.BuildAuthMessage(key)
	case rtsV4:
		msg, err = rts4.BuildAuthMessage(key)
	case rtsV5:
		msg, err = rts5.BuildAuthMessage(key)
	default:
		return nil, errors.New(errInvalidVersion)
	}
	if err != nil {
		return nil, err
	}

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	b, err := v.watch()

	resp := AuthResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New("authorization failed")
	}

	v.state.clientGUID = resp.ClientTokenGUID

	return &resp, err
}

func handleRST3CloudSessionResponse(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	m := msg.GetRtsCloudSessionResponse()

	resp := AuthResponse{
		Status:          AuthStatus(m.StatusCode),
		ClientTokenGUID: m.ClientTokenGuid,
		Success:         m.Success,
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST4CloudSessionResponse(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	m := msg.GetRtsCloudSessionResponse()

	resp := AuthResponse{
		Status:          AuthStatus(m.StatusCode),
		ClientTokenGUID: m.ClientTokenGuid,
		Success:         m.Success,
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST5CloudSessionResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	m := msg.GetRtsCloudSessionResponse()

	resp := AuthResponse{
		Status:          AuthStatus(m.StatusCode),
		ClientTokenGUID: m.ClientTokenGuid,
		Success:         m.Success,
	}

	b, err := resp.Marshal()
	return b, false, err
}
