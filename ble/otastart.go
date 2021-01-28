package ble

import (
	"encoding/json"
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// OTAStartResponse is the unified response for ota start messages
type OTAStartResponse struct {
	Status int
}

// Marshal converts a OTAStartResponse message to bytes
func (sr *OTAStartResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a OTAStartResponse byte slice to a OTAStartResponse
func (sr *OTAStartResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// OTAStart sends a OTAStart message to the vector robot
func (v *VectorBLE) OTAStart(url string) (*OTAStartResponse, error) {
	if !v.state.authorized {
		return nil, errors.New(errNotAuthorized)
	}

	msg, err := rts.BuildOTAStartMessage(v.ble.Version(), url)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	b, err := v.watch()

	resp := OTAStartResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err
}

func handleRSTOtaUpdateResponse(v *VectorBLE, msg interface{}) ([]byte, bool, error) {
	var m *rts.RtsOtaUpdateResponse
	switch v.ble.Version() {

	case rtsV2:
		t, ok := msg.(*rts.RtsConnection_2)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		m = t.GetRtsOtaUpdateResponse()

	case rtsV3:
		t, ok := msg.(*rts.RtsConnection_3)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		m = t.GetRtsOtaUpdateResponse()

	case rtsV4:
		t, ok := msg.(*rts.RtsConnection_4)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		m = t.GetRtsOtaUpdateResponse()

	case rtsV5:
		t, ok := msg.(*rts.RtsConnection_5)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		m = t.GetRtsOtaUpdateResponse()

	default:
		return handlerUnsupportedVersionError()

	}

	resp := OTAStartResponse{
		Status: int(m.Status),
	}

	b, err := resp.Marshal()
	return b, false, err
}
