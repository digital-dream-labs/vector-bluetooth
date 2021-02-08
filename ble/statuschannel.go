package ble

// StatusChannel is a general struct containing status updates for things like OTA + log downloading
type StatusChannel struct {
	LogStatus *statusCounter
	OTAStatus *statusCounter
}

type statusCounter struct {
	PacketNumber uint32
	PacketTotal  uint32
}

func (v *VectorBLE) sendLogStatus(arg *statusCounter) {
	if v.statuschan == nil {
		return
	}
	v.statuschan <- StatusChannel{
		LogStatus: arg,
	}
}

func (v *VectorBLE) sendOTAStatus(arg *statusCounter) {
	if v.statuschan == nil {
		return
	}
	v.statuschan <- StatusChannel{
		OTAStatus: arg,
	}
}
