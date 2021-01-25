package rts5

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildNonceResponse builds the acknowledgement message for the nonce
func BuildNonceResponse() ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_5WithRtsAck(
			&rts.RtsAck{
				RtsConnectionTag: uint8(rts.RtsConnection_5Tag_RtsNonceMessage),
			},
		),
	)
}
