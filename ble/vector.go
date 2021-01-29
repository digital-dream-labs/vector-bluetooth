package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/conn"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

const (
	bleBuffer = 5
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
	filedownload  filedownload
}

type filedownload struct {
	FileID      uint32
	PacketTotal uint32
	// If the logs ever get bigger than a few k, this will
	// definitely have to be rewritten to use a tempfile or something
	Buffer []uint8
}

const (
	errNotAuthorized = "your vector does not have an authorized bluetooth connection"
)

// New returns a new Vector
func New() (*VectorBLE, error) {
	bleReader := make(chan []byte, bleBuffer)

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
