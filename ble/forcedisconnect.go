package ble

func handleRTSForceDisconnect(v *VectorBLE, msg interface{}) (data []byte, cont bool, err error) {
	v.state.authorized = false
	v.state.clientGUID = ""
	v.state.nonceResponse = nil
	return nil, false, nil
}
