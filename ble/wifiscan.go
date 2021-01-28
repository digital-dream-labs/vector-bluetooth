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

// WifiScanResponse is the unified response for wifi scan messages
type WifiScanResponse struct {
	Networks []WifiNetwork `json:"networks,omitempty"`
}

// WifiNetwork is an entry for one network
type WifiNetwork struct {
	WifiSSID       string `json:"wifi_ssid,omitempty"`
	SignalStrength int    `json:"signal_strength,omitempty"`
	AuthType       int    `json:"auth_type,omitempty"`
	Hidden         bool   `json:"hidden,omitempty"`
}

// Marshal converts a WifiScanResponse message to bytes
func (sr *WifiScanResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a WifiScanResponse byte slice to a WifiScanResponse
func (sr *WifiScanResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// WifiScan sends a WifiScan message to the vector robot
func (v *VectorBLE) WifiScan() (*WifiScanResponse, error) {
	if !v.state.authorized {
		return nil, errors.New(errNotAuthorized)
	}

	var (
		msg []byte
		err error
	)

	switch v.ble.Version() {
	case rtsV2:
		msg, err = rts2.BuildWifiScanMessage()
	case rtsV3:
		msg, err = rts3.BuildWifiScanMessage()
	case rtsV4:
		msg, err = rts4.BuildWifiScanMessage()
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

	resp := WifiScanResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err
}

func handleRST2WifiScanResponse(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	m := msg.GetRtsWifiScanResponse2()

	nw := []WifiNetwork{}

	for _, v := range m.ScanResult {
		ssid, _ := hex.DecodeString(v.WifiSsidHex)

		tn := WifiNetwork{
			WifiSSID:       string(ssid),
			SignalStrength: int(v.SignalStrength),
			Hidden:         v.Hidden,
			AuthType:       int(v.AuthType),
		}
		nw = append(nw, tn)
	}

	resp := WifiScanResponse{
		Networks: nw,
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST3WifiScanResponse(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	m := msg.GetRtsWifiScanResponse3()

	nw := []WifiNetwork{}

	for _, v := range m.ScanResult {
		ssid, _ := hex.DecodeString(v.WifiSsidHex)

		tn := WifiNetwork{
			WifiSSID:       string(ssid),
			SignalStrength: int(v.SignalStrength),
			Hidden:         v.Hidden,
			AuthType:       int(v.AuthType),
		}
		nw = append(nw, tn)
	}

	resp := WifiScanResponse{
		Networks: nw,
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST4WifiScanResponse(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	m := msg.GetRtsWifiScanResponse3()

	nw := []WifiNetwork{}

	for _, v := range m.ScanResult {
		ssid, _ := hex.DecodeString(v.WifiSsidHex)

		tn := WifiNetwork{
			WifiSSID:       string(ssid),
			SignalStrength: int(v.SignalStrength),
			Hidden:         v.Hidden,
			AuthType:       int(v.AuthType),
		}
		nw = append(nw, tn)
	}

	resp := WifiScanResponse{
		Networks: nw,
	}

	b, err := resp.Marshal()
	return b, false, err
}

func handleRST5WifiScanResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	m := msg.GetRtsWifiScanResponse3()

	nw := []WifiNetwork{}

	for _, v := range m.ScanResult {
		ssid, _ := hex.DecodeString(v.WifiSsidHex)

		tn := WifiNetwork{
			WifiSSID:       string(ssid),
			SignalStrength: int(v.SignalStrength),
			Hidden:         v.Hidden,
			AuthType:       int(v.AuthType),
		}
		nw = append(nw, tn)
	}

	resp := WifiScanResponse{
		Networks: nw,
	}

	b, err := resp.Marshal()
	return b, false, err
}
