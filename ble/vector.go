package ble

import (
	"fmt"

	"github.com/digital-dream-labs/vector-bluetooth/ble/bleconn"
	"tinygo.org/x/bluetooth"
)

// VectorBLE contains the information required to connect, etc
type VectorBLE struct {
	bleReader chan []byte
	ble       *bleconn.BLEConn
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
func New(name string, adapt *bluetooth.Adapter) (*VectorBLE, error) {
	v := VectorBLE{
		bleReader: make(chan []byte),
	}

	b, err := bleconn.New(v.bleReader)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := b.Connect(name, adapt); err != nil {
			fmt.Println(err)
		}
	}()

	v.ble = b

	return &v, nil
}
