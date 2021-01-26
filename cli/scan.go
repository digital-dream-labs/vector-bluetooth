package cli

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *conf) scan() {
	resp, err := c.v.Scan()
	if err != nil {
		log.Println("unable to scan: ", err)
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to scan: ", err)
	}

	fmt.Println(string(data))
	fmt.Println("To connect, type \"connect ID\"")
}
