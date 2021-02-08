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
	v      *ble.VectorBLE
	status chan ble.StatusChannel
}

// BLEShell starts the bluetooth interactive shell
func BLEShell() {
	sc := make(chan ble.StatusChannel)
	v, err := ble.New(
		ble.WithLogDirectory("."),
		ble.WithStatusChan(sc),
	)
	if err != nil {
		log.Fatal(err)
	}
	c := conf{
		v:      v,
		status: sc,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("type \"help\" for a list of commands")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		args := strings.Fields(
			strings.ReplaceAll(text, "\n", ""),
		)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "scan":
			c.scan()
		case "connect":
			c.vectorConnect(args)
		case "authorize":
			c.auth(args)
		case "configure":
			c.configure()
		case "get-status":
			c.getStatus()
		case "wifi-scan":
			c.wifiScan()
		case "wifi-connect":
			c.wifiConnect(args)
		case "wifi-ip":
			c.wifiIP()
		case "wifi-forget":
			c.wifiForget(args)
		case "ota-start":
			c.startOTA(args)
		case "ota-cancel":
			c.cancelOTA()
		case "logs":
			c.logs()
		default:
			help()
		}
	}
}
