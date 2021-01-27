package conn

import (
	"github.com/currantlabs/ble"
)

func discoverFilter() ble.AdvFilter {
	return func(a ble.Advertisement) bool {
		for _, s := range a.Services() {
			if s.String() == "fee3" {
				return true
			}
		}
		return false
	}
}

func deviceFilter(device ble.Addr) ble.AdvFilter {
	return func(a ble.Advertisement) bool {
		for _, s := range a.Services() {
			if s.String() == "fee3" && a.Address().String() == device.String() {
				return true
			}
		}
		return false
	}
}
