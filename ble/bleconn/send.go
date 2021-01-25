package bleconn

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

const (
	maxSize = 20
)

// Send sends a message via BLE
func (b *BLEConn) Send(buf []byte) error {
	if b.encrypted {
		if b.encrypted {
			var err error
			buf, err = b.Crypto.Encrypt(buf)
			if err != nil {
				fmt.Println("decrypt error: ", err)
				return err
			}
		}

	}
	// fmt.Println("msg is ", buf)

	size := len(buf)
	if size < maxSize {
		return sendRawmessage(b.readchar, msgSolo, buf, size)
	}

	remaining := len(buf)

	for remaining > 0 {
		offset := size - remaining

		switch {

		case remaining == size:
			msgSize := maxSize - 1
			if err := sendRawmessage(b.readchar, msgStart, buf[offset:offset+msgSize], msgSize); err != nil {
				return err
			}
			remaining -= msgSize

		case remaining < maxSize:
			if err := sendRawmessage(b.readchar, msgEnd, buf[offset:], remaining); err != nil {
				return err
			}
			remaining = 0

		default:
			fmt.Println("middle msg")
			msgSize := maxSize - 1
			if err := sendRawmessage(b.readchar, msgContinue, buf[offset:offset+msgSize], msgSize); err != nil {
				return err
			}
			remaining -= msgSize
		}

	}
	return nil
}

func sendRawmessage(char bluetooth.DeviceCharacteristic, multipart byte, buffer []byte, size int) error {
	var msg []byte
	msg = append(msg, getHeaderByte(multipart, size))
	msg = append(msg, buffer...)
	_, err := char.WriteWithoutResponse(msg)
	return err
}

func getHeaderByte(multipart byte, size int) byte {
	//nolint
	return byte(((int(multipart) << 6) | (size & -193)))
}
