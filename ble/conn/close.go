package conn

// Close shuts down the BLE connection
func (c *Connection) Close() error {
	return c.device.Stop()
}

// Reset clears all connection information
func (c *Connection) Reset() {
	c.connected = false
	c.encrypted = false
}
