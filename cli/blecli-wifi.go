package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/digital-dream-labs/vector-bluetooth/ble"
)

func wifiScan(v *ble.VectorBLE) {
	resp, err := v.WifiScan()
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	fmt.Println(string(data))

}

func wifiConnect(v *ble.VectorBLE, args []string) {
	//nolint
	if len(args) != 4 {
		fmt.Println("invalid argument.  Usage is \n wifi-connect SSID PASSWORD NETWORKTYPE")
	}

	t, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("invalid argument.  Usage is \n wifi-connect SSID PASSWORD NETWORKTYPE")
	}

	resp, err := v.WifiConnect(args[1], args[2], 10, t)
	if err != nil {
		log.Println("unable to get status: ", err)
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	fmt.Println(string(data))

}
