package rts3

import (
	"bytes"

	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

func buildMessage(message *rts.RtsConnection_3) ([]byte, error) {
	ec := rts.NewExternalCommsWithRtsConnection(
		rts.NewRtsConnectionWithRtsConnection3(
			message,
		),
	)
	var br bytes.Buffer
	if err := ec.Pack(&br); err != nil {
		return nil, err
	}
	return br.Bytes(), nil
}
