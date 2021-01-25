package bleconn

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

// Connect connects to the ble device
func (b *BLEConn) Connect(name string, adapter *bluetooth.Adapter) error {
	/*
		var cln ble.Client

		ctx := ble.WithSigHandler(
			context.WithTimeout(
				context.Background(),
				scanDuration*2,
			),
		)

		fmt.Printf("connecting to %s\n", name)
		cln, err := ble.Connect(
			ctx,
			deviceFilter(ble.NewAddr(name)),
		)
		if err != nil {
			// return err
			fmt.Println(err)
		}

		b.conn.Clients[name] = cln

		p, err := cln.DiscoverProfile(false)
		if err != nil {
			return errors.Wrap(err, "can't discover profile")
		}

		b.conn.profile = p

		return nil
	*/

	// FIXME:  Right now, if there's any error in anything, everything blows up.  Make
	// an error channel and handle it appropriately

	var dev *bluetooth.Device
	var err error
	go func() {
		_ = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
			if device.LocalName() == name {
				dev, err = adapter.Connect(
					device.Address,
					bluetooth.ConnectionParams{
						// ADD A TIMEOUT
					},
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				b.device = dev
				if err = b.getDataChannels(); err != nil {
					fmt.Println("data channel error: ", err)
					return
				}
				buf := bleBuffer{}

				//nolint
				for {
					select {
					case incoming := <-b.write:
						buf := buf.receiveRawBuffer(incoming)
						if buf == nil {
							continue
						}

						switch {
						case !b.connected:
							// FIXME:  on reconnect, this will explode.
							fmt.Println("connection request")
							_ = b.connRequest(incoming)
							b.connected = true
							b.Version = int(incoming[2])
							continue
						default:
							if b.encrypted {
								var err error
								buf, err = b.Crypto.DecryptMessage(buf)
								if err != nil {
									fmt.Println("decrypt error: ", err)
								}
							}
							b.out <- buf
						}

					}
				}

			}
		})
	}()

	return nil
}

func (b *BLEConn) connRequest(buf []byte) error {
	_, err := b.readchar.WriteWithoutResponse(buf)
	if err != nil {
		fmt.Println("send error: ", err)
	}
	return nil
}

/*
func deviceFilter(device ble.Addr) ble.AdvFilter {
	return func(a ble.Advertisement) bool {
		fmt.Println(a.LocalName())
		for _, s := range a.Services() {
			fmt.Println(s)
			fmt.Println(a.Address())
			if s.String() == "fee3" && a.Address().String() == device.String() {
				fmt.Println("true for ", a.LocalName())
				return true
			}
		}
		return false
	}
}
*/
