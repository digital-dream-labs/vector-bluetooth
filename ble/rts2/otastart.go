package rts2

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildOTAStartMessage builds the ota start message
func BuildOTAStartMessage(url string) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_2WithRtsOtaUpdateRequest(
			&rts.RtsOtaUpdateRequest{
				Url: url,
			},
		),
	)
}
