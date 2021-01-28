package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

func handleRTS2NonceRequest(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	if err := v.ble.SetNonces(msg.GetRtsNonceMessage()); err != nil {
		return nil, false, err
	}

	b, err := rts2.BuildNonceResponse()
	if err != nil {
		return nil, false, err
	}

	v.state.nonceResponse = b

	return nil, false, nil
}

func handleRTS3NonceRequest(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	if err := v.ble.SetNonces(msg.GetRtsNonceMessage()); err != nil {
		return nil, false, err
	}

	b, err := rts3.BuildNonceResponse()
	if err != nil {
		return nil, false, err
	}

	v.state.nonceResponse = b

	return nil, false, nil
}

func handleRTS4NonceRequest(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	if err := v.ble.SetNonces(msg.GetRtsNonceMessage()); err != nil {
		return nil, false, err
	}

	b, err := rts4.BuildNonceResponse()
	if err != nil {
		return nil, false, err
	}

	v.state.nonceResponse = b

	return nil, false, nil
}

func handleRTS5NonceRequest(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	if err := v.ble.SetNonces(msg.GetRtsNonceMessage()); err != nil {
		return nil, false, err
	}

	b, err := rts5.BuildNonceResponse()
	if err != nil {
		return nil, false, err
	}

	v.state.nonceResponse = b

	return nil, false, nil
}
