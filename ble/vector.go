package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/conn"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// VectorBLE contains the information required to connect, etc
type VectorBLE struct {
	bleReader chan []byte
	ble       bleconn
	state     state
}

type bleconn interface {
	Connect(int) error
	Connected() bool
	Close() error
	EnableEncryption()
	GetRemotePublicKey() [32]uint8
	Scan() (*conn.ScanResponse, error)
	SetNonces(msg *rts.RtsNonceMessage) error
	SetRemotePublicKey(msg *rts.RtsConnRequest) error
	Send(buf []byte) error
	SetPin(string) error
	Version() int
}

type state struct {
	nonceResponse []byte
	authorized    bool
	clientGUID    string
}

const (
	errNotAuthorized  = "your vector does not have an authorized bluetooth connection"
	errInvalidVersion = "invalid rts version"
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
