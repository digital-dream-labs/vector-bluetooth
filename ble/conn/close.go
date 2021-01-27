package conn

// Close shuts down the BLE connection
func (c *Connection) Close() error {
	return c.device.Stop()
}
