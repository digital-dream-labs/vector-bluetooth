package rts3

import (
	"encoding/hex"

	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildWifiConnectMessage builds the wifi connect message
func BuildWifiConnectMessage(ssid, password string, timeout, authtype int) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_3WithRtsWifiConnectRequest(
			&rts.RtsWifiConnectRequest{
				WifiSsidHex: hex.EncodeToString([]byte(ssid)),
				Password:    password,
				Timeout:     uint8(timeout),
				AuthType:    uint8(authtype),
			},
		),
	)
}
