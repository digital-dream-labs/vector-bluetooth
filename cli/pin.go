package cli

import (
	"fmt"
	"log"
)

func (c *conf) setPin(args []string) {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
	}

	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n set-pin XXXXXX")
		return
	}

	err := c.v.SendPin(args[1])
	if err != nil {
		log.Println("unable to set pin: ", err)
		return
	}

}
