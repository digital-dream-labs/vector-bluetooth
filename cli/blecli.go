package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digital-dream-labs/vector-bluetooth/ble"
)

type conf struct {
	v *ble.VectorBLE
}

// BLEShell starts the bluetooth interactive shell
func BLEShell() {
	// bkrt := "Vector G4T1"

	v, err := ble.New()
	if err != nil {
		log.Fatal(err)
	}
	c := conf{
		v: v,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("type \"help\" for a list of commands")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		args := strings.Fields(
			strings.ReplaceAll(text, "\n", ""),
		)

		switch args[0] {
		case "scan":
			c.scan()
		case "connect":
			c.vectorConnect(args)
		case "get-status":
			c.getStatus()
		case "wifi-scan":
			c.wifiScan()
		case "wifi-connect":
			c.wifiConnect(args)
		case "ota-start":
			c.startOTA(args)
		default:
			help()
		}

	}
}
