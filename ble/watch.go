package ble

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

func (v *VectorBLE) watch() ([]byte, error) {
	var (
		resp []byte
		cont bool
		err  error
	)

	// cont tells the loop whether to continue watching or not.
	cont = true

	for {
		if cont {
			incoming := <-v.bleReader
			z := bytes.NewBuffer(incoming)
			comm := rts.ExternalComms{}
			if err := comm.Unpack(z); err != nil {
				return nil, err
			}
			m := comm.GetRtsConnection()

			if m == nil {
				return nil, errors.New("empty rts connection")
			}

			switch m.Tag() {

			case rts.RtsConnectionTag_Error:

			case rts.RtsConnectionTag_RtsConnection2:
				f, ok := rts2Handlers[m.GetRtsConnection2().Tag()]
				if !ok {
					cont = false
					err = fmt.Errorf("unsupported message: %v", m.GetRtsConnection2().Tag())
				}
				resp, cont, err = f(v, m.GetRtsConnection2())

			case rts.RtsConnectionTag_RtsConnection3:
				f, ok := rts3Handlers[m.GetRtsConnection3().Tag()]
				if !ok {
					cont = false
					err = fmt.Errorf("unsupported message: %v", m.GetRtsConnection3().Tag())
				}
				resp, cont, err = f(v, m.GetRtsConnection3())

			case rts.RtsConnectionTag_RtsConnection4:
				f, ok := rts4Handlers[m.GetRtsConnection4().Tag()]
				if !ok {
					cont = false
					err = fmt.Errorf("unsupported message: %v", m.GetRtsConnection4().Tag())
				}
				resp, cont, err = f(v, m.GetRtsConnection4())

			case rts.RtsConnectionTag_RtsConnection5:
				f, ok := rts5Handlers[m.GetRtsConnection5().Tag()]
				if !ok {
					cont = false
					err = fmt.Errorf("unsupported message: %v", m.GetRtsConnection5().Tag())
				}
				resp, cont, err = f(v, m.GetRtsConnection5())

			case rts.RtsConnectionTag_INVALID:
				cont = false
				err = errors.New("invalid message")

			default:
				cont = false
				err = errors.New("unsupported message version")
			}

		} else {
			return resp, err
		}
	}
}
