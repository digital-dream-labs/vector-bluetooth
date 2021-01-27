package ble

import "github.com/digital-dream-labs/vector-bluetooth/ble/conn"

// Scan shows a list of connectable vectors
func (v *VectorBLE) Scan() (*conn.ScanResponse, error) {
	return v.ble.Scan()
}

// Connected returns true if the device is connected
func (v *VectorBLE) Connected() bool {
	return v.ble.Connected()
}

// Close stops the BLE connection
func (v *VectorBLE) Close() error {
	return v.ble.Close()
}
