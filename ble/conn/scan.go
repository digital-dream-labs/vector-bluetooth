package conn

import (
	"context"
	"time"

	"github.com/currantlabs/ble"
)

const (
	scanDuration = 5 * time.Second
)

// ScanResponse is a list of devices the BLE adaptor has found
type ScanResponse struct {
	Devices []*Device
}

// Device is a single device entity
type Device struct {
	ID      int
	Name    string
	Address string
}

// Scan looks for BLE devices matching the vector requirements
func (c *Connection) Scan() (*ScanResponse, error) {
	ctx := ble.WithSigHandler(
		context.WithTimeout(
			context.Background(),
			scanDuration,
		),
	)

	h := newADVHandler(c)

	_ = ble.Scan(
		ctx,
		false,
		h.scan,
		discoverFilter(),
	)

	d := []*Device{}

	for k, v := range c.scanresults {
		td := Device{
			ID:      k,
			Name:    v.name,
			Address: v.addr.String(),
		}
		d = append(d, &td)

	}

	resp := ScanResponse{
		Devices: d,
	}

	return &resp, nil

}

type advhandler struct {
	count int
	conn  *Connection
}

func newADVHandler(conn *Connection) *advhandler {
	return &advhandler{
		count: 1,
		conn:  conn,
	}
}

func (a *advhandler) scan(d ble.Advertisement) {
	if d.Connectable() {

		if a.conn.scanresults != nil {
			for _, v := range a.conn.scanresults {
				if v.name == d.LocalName() {
					return
				}
			}
		}

		a.conn.scanresults[a.count] = scanresult{
			name: d.LocalName(),
			addr: d.Address(),
		}
		a.count++
	}

}
