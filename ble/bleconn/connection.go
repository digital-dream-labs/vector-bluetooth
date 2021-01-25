package bleconn

/*
import (
	"fmt"

	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"
	"github.com/currantlabs/ble/linux"
	"github.com/pkg/errors"
)

// Connection is the configuration struct for ble connections
type Connection struct {
	device  ble.Device
	Devices []string
	// client  ble.Client
	Clients   map[string]ble.Client
	readUUID  ble.UUID
	writeUUID ble.UUID
	addr      ble.Addr
	profile   *ble.Profile
}

//// NEW

func bleConnect() (*Connection, error) {
	r, err := ble.Parse(readUUID)
	if err != nil {
		return nil, errors.New("cannot parse read uuid")
	}

	w, err := ble.Parse(writeUUID)
	if err != nil {
		return nil, errors.New("cannot parse write uuid")
	}

	c := Connection{
		readUUID:  r,
		writeUUID: w,
		Clients:   make(map[string]ble.Client),
	}

	fmt.Printf("Initializing device ...\n")

	d, err := dev.NewDevice("default")
	if err != nil {
		return nil, errors.Wrap(err, "can't new device")
	}
	ble.SetDefaultDevice(d)
	c.device = d

	dev, ok := d.(*linux.Device)
	if !ok {
		return nil, errors.New("not a device")
	}

	if err := updateLinuxParam(dev); err != nil {
		return nil, errors.Wrap(updateLinuxParam(dev), "can't update hci parameters")
	}

	return &c, nil
}
*/
