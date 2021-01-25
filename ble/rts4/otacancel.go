package rts4

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildOTACancelMessage builds the ota cancel message
func BuildOTACancelMessage() ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_4WithRtsOtaCancelRequest(
			&rts.RtsOtaCancelRequest{},
		),
	)
}
