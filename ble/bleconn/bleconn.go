package bleconn

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/blecrypto"
	"tinygo.org/x/bluetooth"
)

/*
const (
	scanDuration = 5 * time.Second
	readUUID     = "7d2a4bda-d29b-4152-b725-2491478c5cd7"
	writeUUID    = "30619f2d-0f54-41bd-a65a-7588d8c85b45"
)
*/

// BLEConn is a ble connection to a vector
type BLEConn struct {
	device    *bluetooth.Device
	readchar  bluetooth.DeviceCharacteristic
	writechar bluetooth.DeviceCharacteristic
	Crypto    *blecrypto.BLECrypto
	Version   int
	// rts       rtsHandler
	read      chan []byte
	write     chan []byte
	out       chan []byte
	encrypted bool
	connected bool
}

// New returns a new BLEConn
func New(output chan []byte) (*BLEConn, error) {
	read := make(chan []byte)
	write := make(chan []byte)

	v := BLEConn{
		read:   read,
		write:  write,
		out:    output,
		Crypto: blecrypto.New(),
	}

	return &v, nil
}

// EnableEncryption sets the encryption bit to automatically de/encrypt
func (b *BLEConn) EnableEncryption() {
	b.encrypted = true
}
