package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/conn"
)

// VectorBLE contains the information required to connect, etc
type VectorBLE struct {
	bleReader chan []byte
	ble       *conn.Connection
	state     state
}

type state struct {
	nonceResponse []byte
	authorized    bool
	clientGUID    string
}

const (
	errNotAuthorized = "your vector does not have an authorized bluetooth connection"
)

// New returns a new Vector
func New() (*VectorBLE, error) {
	bleReader := make(chan []byte)

	b, err := conn.New(bleReader)
	if err != nil {
		return nil, err
	}
	v := VectorBLE{
		bleReader: bleReader,
		ble:       b,
	}

	return &v, nil
}
