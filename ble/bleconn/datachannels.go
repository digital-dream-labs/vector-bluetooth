package bleconn

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

const (
	readCharService  = "7d2a4bda-d29b-4152-b725-2491478c5cd7"
	writeCharService = "30619f2d-0f54-41bd-a65a-7588d8c85b45"
	service          = "0000fee3-0000-1000-8000-00805f9b34fb"
)

func (b *BLEConn) getDataChannels() error {
	svc, _ := bluetooth.ParseUUID(service)
	read, _ := bluetooth.ParseUUID(readCharService)
	write, _ := bluetooth.ParseUUID(writeCharService)

	srvcs, err := b.device.DiscoverServices(
		[]bluetooth.UUID{svc},
	)
	if err != nil {
		return err
	}

	if len(srvcs) == 0 {
		return fmt.Errorf("no services found")
	}

	chars, err := srvcs[0].DiscoverCharacteristics(
		[]bluetooth.UUID{read, write},
	)
	if err != nil {
		return err
	}

	for _, char := range chars {
		switch char.UUID().String() {

		case readCharService:
			b.readchar = char
			go func() {
				_ = char.EnableNotifications(func(buf []byte) {
					b.read <- buf
				})
			}()

		case writeCharService:
			b.writechar = char
			go func() {
				_ = char.EnableNotifications(func(buf []byte) {
					b.write <- buf
				})
			}()

		}
	}

	return nil
}
