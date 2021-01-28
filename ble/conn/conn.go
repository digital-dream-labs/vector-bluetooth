package conn

import (
	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"
	"github.com/digital-dream-labs/vector-bluetooth/ble/blecrypto"
	"github.com/pkg/errors"
)

// Connection is the configuration struct for ble connections
type Connection struct {
	device      ble.Device
	scanresults map[int]scanresult
	connection  ble.Client
	reader      *ble.Characteristic
	writer      *ble.Characteristic
	profile     *ble.Profile
	incoming    chan []byte
	out         chan []byte
	crypto      *blecrypto.BLECrypto
	version     int
	established bool
	connected   bool
	encrypted   bool
}

type scanresult struct {
	name string
	addr ble.Addr
}

// New returns a connection, or an error on failure
func New(output chan []byte) (*Connection, error) {
	c := Connection{
		scanresults: make(map[int]scanresult),
		incoming:    make(chan []byte),
		out:         output,
		crypto:      blecrypto.New(),
	}

	d, err := dev.NewDevice("default")
	if err != nil {
		return nil, errors.Wrap(err, "can't add new device")
	}
	ble.SetDefaultDevice(d)
	c.device = d

	return &c, nil
}

// EnableEncryption sets the encryption bit to automatically de/encrypt
func (c *Connection) EnableEncryption() {
	c.encrypted = true
}

// Connected lets external packages know if the initial connection attempt has happened
func (c *Connection) Connected() bool {
	return c.connected
}
