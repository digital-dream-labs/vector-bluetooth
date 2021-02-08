package cli

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *conf) logs() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
	}
	go c.renderLogStatus()

	resp, err := c.v.DownloadLogs()
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
