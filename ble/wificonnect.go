package ble

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// WifiConnectResponse is the unified response for wifi connect messages
type WifiConnectResponse struct {
	WifiSSID string `json:"wifi_ssid,omitempty"`
	State    int    `json:"state,omitempty"`
	Result   int    `json:"result,omitempty"`
}

// Marshal converts a WifiConnectResponse message to bytes
func (sr *WifiConnectResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a WifiConnectResponse byte slice to a WifiConnectResponse
func (sr *WifiConnectResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// WifiConnect sends a wifi connect message to the robot
func (v *VectorBLE) WifiConnect(ssid string, password string, timeout int, authtype int) (*WifiConnectResponse, error) {
	if !v.state.authorized {
		return nil, errors.New(errNotAuthorized)
	}

	var (
		msg []byte
		err error
	)

	switch v.ble.Version() {
	case rtsV2:
		msg, err = rts2.BuildWifiConnectMessage(ssid, password, timeout, authtype)
	case rtsV3:
		msg, err = rts3.BuildWifiConnectMessage(ssid, password, timeout, authtype)
	case rtsV4:
		msg, err = rts4.BuildWifiConnectMessage(ssid, password, timeout, authtype)
	case rtsV5:
		msg, err = rts5.BuildWifiConnectMessage(ssid, password, timeout, authtype)
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

	resp := WifiConnectResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err

}

func handleRST2WifiConnectionResponse(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	m := msg.GetRtsWifiConnectResponse()

	ssid, _ := hex.DecodeString(m.WifiSsidHex)

	resp := WifiConnectResponse{
		WifiSSID: string(ssid),
		State:    int(m.WifiState),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST3WifiConnectionResponse(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	m := msg.GetRtsWifiConnectResponse3()

	ssid, _ := hex.DecodeString(m.WifiSsidHex)

	resp := WifiConnectResponse{
		WifiSSID: string(ssid),
		State:    int(m.WifiState),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST4WifiConnectionResponse(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	m := msg.GetRtsWifiConnectResponse3()

	ssid, _ := hex.DecodeString(m.WifiSsidHex)

	resp := WifiConnectResponse{
		WifiSSID: string(ssid),
		State:    int(m.WifiState),
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST5WifiConnectionResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	m := msg.GetRtsWifiConnectResponse3()

	ssid, _ := hex.DecodeString(m.WifiSsidHex)

	resp := WifiConnectResponse{
		WifiSSID: string(ssid),
		State:    int(m.WifiState),
		Result:   int(m.ConnectResult),
	}

	b, err := resp.Marshal()
	return b, false, err
}
