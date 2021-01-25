package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/digital-dream-labs/vector-bluetooth/ble"
)

func getStatus(v *ble.VectorBLE) {
	resp, err := v.GetStatus()
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("unable to get status: ", err)
	}

	fmt.Println(string(data))

}
