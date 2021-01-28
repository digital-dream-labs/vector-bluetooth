package ble

import "errors"

func handlerUnsupportedVersionError() ([]byte, bool, error) {
	return nil, false, errors.New("unsupported rts protocol version")
}

func handlerUnsupportedTypeError() ([]byte, bool, error) {
	return nil, false, errors.New("unsupported message type")
}
