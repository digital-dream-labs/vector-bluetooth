package ble

func handleRTSForceDisconnect(v *VectorBLE, msg interface{}) ([]byte, bool, error) {
	v.state.authorized = false
	v.state.clientGUID = ""
	v.state.nonceResponse = nil
	return nil, false, nil
}
