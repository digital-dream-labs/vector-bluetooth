package conn

import (
	"github.com/currantlabs/ble"
)

const vectorservice = "fee3"

func discoverFilter() ble.AdvFilter {
	return func(a ble.Advertisement) bool {
		for _, s := range a.Services() {
			if s.String() == vectorservice {
				return true
			}
		}
		return false
	}
}

func deviceFilter(device ble.Addr) ble.AdvFilter {
	return func(a ble.Advertisement) bool {
		for _, s := range a.Services() {
			if s.String() == vectorservice && a.Address().String() == device.String() {
				return true
			}
		}
		return false
	}
}
