package rts3

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildStatusMessage builds the status request message
func BuildStatusMessage() ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_3WithRtsStatusRequest(
			&rts.RtsStatusRequest{},
		),
	)
}
