package conn

import (
	"github.com/go-ble/ble"
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
			if s.String() == "fee3" && a.Addr().String() == device.String() {
				return true
			}
		}
		return false
	}
}
