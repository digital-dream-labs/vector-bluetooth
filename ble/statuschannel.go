package ble

// StatusChannel is a general struct containing status updates for things like OTA + log downloading
type StatusChannel struct {
	LogStatus *StatusCounter
	OTAStatus *StatusCounter
}

type StatusCounter struct {
	PacketNumber uint32
	PacketTotal  uint32
	Error        string
}

func (v *VectorBLE) sendLogStatus(arg *StatusCounter) {
	if v.statuschan == nil {
		return
	}
	v.statuschan <- StatusChannel{
		LogStatus: arg,
	}
}

func (v *VectorBLE) sendOTAStatus(arg *StatusCounter) {
	if v.statuschan == nil {
		return
	}
	v.statuschan <- StatusChannel{
		OTAStatus: arg,
	}
}
