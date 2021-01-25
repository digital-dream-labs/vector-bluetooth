package ble

import (
	"fmt"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

func handleRts2ConnRequest(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	if err := v.ble.Crypto.SetRemotePublicKey(msg.GetRtsConnRequest()); err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	b, err := rts2.GetConnResponse(v.ble.Crypto.GetRemotePublicKey())
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}

func handleRts5ConnRequest(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	if err := v.ble.Crypto.SetRemotePublicKey(msg.GetRtsConnRequest()); err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	b, err := rts5.GetConnResponse(v.ble.Crypto.GetRemotePublicKey())
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}
