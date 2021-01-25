package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digital-dream-labs/vector-bluetooth/ble"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

// BLEShell starts the bluetooth interactive shell
func BLEShell() {
	// Enable BLE interface.
	_ = adapter.Enable()

	// bkrt := "Vector G4T1"

	var bkrt string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the name of your robot, ie \"Vector G4T1\"")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		bkrt = strings.ReplaceAll(text, "\n", "")
		break
	}

	v, err := ble.New(bkrt, adapter)
	if err != nil {
		log.Fatal(err)
	}

	if err := v.SignOn(); err != nil {
		log.Fatal(err)
	}

	var key string
	fmt.Println("Please enter the key from the screen of your vector")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		key = strings.ReplaceAll(text, "\n", "")
		break
	}

	if err := v.SendPin(key); err != nil {
		log.Fatal(err)
	}

	fmt.Println("you are now authorized!  type \"help\" for a list of commands")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		args := strings.Fields(
			strings.ReplaceAll(text, "\n", ""),
		)

		switch args[0] {
		case "get-status":
			getStatus(v)
		case "wifi-scan":
			wifiScan(v)
		case "wifi-connect":
			wifiConnect(v, args)
		case "ota-start":
			startOTA(v, args)
		}
	}

}
