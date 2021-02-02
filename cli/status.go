package cli

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *conf) getStatus() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
	}

	resp, err := c.v.GetStatus()
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
