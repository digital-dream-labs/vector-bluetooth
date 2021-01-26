package cli

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *conf) auth(args []string) {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n auth TOKEN")
		return
	}

	resp, err := c.v.Auth(args[1])
	if err != nil {
		log.Println("authorization call error: ", err)
		return
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
		return
	}

	fmt.Println(string(data))
}
