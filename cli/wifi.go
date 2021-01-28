package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (c *conf) wifiScan() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	resp, err := c.v.WifiScan()
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	fmt.Println(string(data))
}

func (c *conf) wifiConnect(args []string) {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	//nolint
	if len(args) != 4 {
		fmt.Println("invalid argument.  Usage is \n wifi-connect SSID PASSWORD NETWORKTYPE")
		return
	}

	t, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("invalid argument.  Usage is \n wifi-connect SSID PASSWORD NETWORKTYPE")
		return
	}

	resp, err := c.v.WifiConnect(args[1], args[2], 10, t)
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	fmt.Println(string(data))

}

func (c *conf) wifiIP() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	resp, err := c.v.WifiIP()
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	fmt.Println(string(data))
}

func (c *conf) wifiForget(args []string) {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n wifi-connect SSID PASSWORD NETWORKTYPE")
		return
	}

	resp, err := c.v.WifiForget(args[1])
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	fmt.Println(string(data))

}
