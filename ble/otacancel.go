package ble

import (
	"errors"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
)

// OTACancel sends a OTACancel message to the vector robot
func (v *VectorBLE) OTACancel() ([]byte, error) {
	if !v.state.authorized {
		return nil, errors.New(errNotAuthorized)
	}

	var (
		msg []byte
		err error
	)

	switch v.ble.Version() {
	case rtsV2:
		msg, err = rts2.BuildOTACancelMessage()
	case rtsV3:
		msg, err = rts3.BuildOTACancelMessage()
	case rtsV4:
		msg, err = rts4.BuildOTACancelMessage()
	case rtsV5:
		msg, err = rts5.BuildOTACancelMessage()
	default:
		return nil, errors.New(errInvalidVersion)
	}
	if err != nil {
		return nil, err
	}

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	_, err = v.watch()

	return nil, err
}
