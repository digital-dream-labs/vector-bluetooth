package cli

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *conf) startOTA(args []string) {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
	}

	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n ota-start URL")
		return
	}

	go c.renderOTAStatus()

	resp, err := c.v.OTAStart(args[1])
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

func (c *conf) cancelOTA() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
	}
	resp, err := c.v.OTACancel()
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
