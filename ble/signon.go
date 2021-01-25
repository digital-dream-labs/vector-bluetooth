package ble

// SignOn initiates the sign-on process
func (v *VectorBLE) SignOn() error {
	_, err := v.watch()
	return err
}
