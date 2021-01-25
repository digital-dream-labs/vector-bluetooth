package rts2

import "github.com/digital-dream-labs/vector-bluetooth/rts"

// BuildWifiScanMessage builds the wifi scan message
func BuildWifiScanMessage() ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_2WithRtsWifiScanRequest(
			&rts.RtsWifiScanRequest{},
		),
	)
}
