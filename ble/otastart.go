package ble

import (
	"encoding/json"
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
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

	var (
		msg []byte
		err error
	)

	switch v.ble.Version() {
	case rtsV2:
		msg, err = rts2.BuildOTAStartMessage(url)
	case rtsV3:
		msg, err = rts3.BuildOTAStartMessage(url)
	case rtsV4:
		msg, err = rts4.BuildOTAStartMessage(url)
	case rtsV5:
		msg, err = rts5.BuildWifiScanMessage()
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

	resp := OTAStartResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err
}

func handleRST2OtaUpdateResponse(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	m := msg.GetRtsOtaUpdateResponse()

	resp := OTAStartResponse{
		Status: int(m.Status),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST3OtaUpdateResponse(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	m := msg.GetRtsOtaUpdateResponse()

	resp := OTAStartResponse{
		Status: int(m.Status),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST4OtaUpdateResponse(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	m := msg.GetRtsOtaUpdateResponse()

	resp := OTAStartResponse{
		Status: int(m.Status),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST5OtaUpdateResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	m := msg.GetRtsOtaUpdateResponse()

	resp := OTAStartResponse{
		Status: int(m.Status),
	}

	b, err := resp.Marshal()
	return b, false, err
}
