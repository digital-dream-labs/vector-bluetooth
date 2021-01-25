package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
)

// OTACancel sends a OTACancel message to the vector robot
func (v *VectorBLE) OTACancel() ([]byte, error) {
	var (
		msg []byte
		err error
	)

	switch v.ble.Version {
	case rtsV2:
		msg, err = rts2.BuildOTACancelMessage()
	case rtsV3:
	case rtsV4:
	case rtsV5:
		msg, err = rts5.BuildOTACancelMessage()
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
