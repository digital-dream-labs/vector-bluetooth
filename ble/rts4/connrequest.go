package rts4

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// GetConnResponse builds the RTS5 connection response
func GetConnResponse(pubkey [32]uint8) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_4WithRtsConnResponse(
			&rts.RtsConnResponse{
				ConnectionType: rts.RtsConnType_FirstTimePair,
				PublicKey:      pubkey,
			},
		),
	)
}
