package cli

import "fmt"

func help() {
	fmt.Println(`vector-cli bluetooth module

	This utility is meant to perform BLE operations to your vector.  The following commands are valid:

	scan                   runs a BLE scan and displays an appropriate list of devices
	connect                connect to a vector via ID (displayed in the scan)
	authorize              performs a cloud authorization (but you'll need to find your token!)
	configure              allows you to make/change configuration
	get-status             displays the status of your vector
	wifi-scan              scan for a list of available wifi networks
	wifi-connect           connect to a wifi network
	ota-start              perform an OTA code download`)
}
