package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/digital-dream-labs/vector-bluetooth/ble"
)

func startOTA(v *ble.VectorBLE, args []string) {
	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n ota-start URL")
	}

	resp, err := v.OTAStart(args[1])
	if err != nil {
		log.Println("unable to get status: ", err)
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	fmt.Println(string(data))

}

func cancelOTA(v *ble.VectorBLE) {
	resp, err := v.OTACancel()
	if err != nil {
		log.Println("unable to get status: ", err)
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	fmt.Println(string(data))
}
