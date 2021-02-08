package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/digital-dream-labs/vector-bluetooth/ble"
)

func (c *conf) configure() {
	if !c.v.Connected() {
		fmt.Println("bluetooth connectivity must be established to use this command")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	v := ble.VectorSettings{}

	for {
		fmt.Print("Enter your timezone: ")
		text, _ := reader.ReadString('\n')
		v.Timezone = strings.ReplaceAll(text, "\n", "")
		break
	}

	for {
		fmt.Print("Enter your default location (ie: San Francisco, California, United States): ")
		text, _ := reader.ReadString('\n')
		v.DefaultLocation = strings.ReplaceAll(text, "\n", "")
		break
	}

	for {
		fmt.Print("Enter your locale: ")
		text, _ := reader.ReadString('\n')
		v.Locale = strings.ReplaceAll(text, "\n", "")
		break
	}

	for {
		fmt.Print("Would you like to send data analytics to DDL (true|false): ")
		text, _ := reader.ReadString('\n')
		b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
		if err != nil {
			fmt.Println(err)
			continue
		}
		v.AllowDataAnalytics = b
		break
	}

	for {
		fmt.Print("Would you like distance to be measured in metric units (true|false): ")
		text, _ := reader.ReadString('\n')
		b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
		if err != nil {
			continue
		}
		v.MetricDistance = b
		break
	}

	for {
		fmt.Print("Would you like temperature to be measured in metric units (true|false): ")
		text, _ := reader.ReadString('\n')
		b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
		if err != nil {
			continue
		}
		v.MetricTemperature = b
		break
	}

	for {
		fmt.Print("Would you like to enable alexa support? (true|false): ")
		text, _ := reader.ReadString('\n')
		b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
		if err != nil {
			continue
		}
		v.AlexaOptIn = b
		break
	}

	if v.AlexaOptIn {
		for {
			fmt.Print("Would you like the button to activate alexa? (true|false): ")
			text, _ := reader.ReadString('\n')
			b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
			if err != nil {
				continue
			}
			v.ButtonWakeword = b
			break
		}
	}

	for {
		fmt.Print("Would you like your clock to be displayed in 24 hour format? (true|false): ")
		text, _ := reader.ReadString('\n')
		b, err := strconv.ParseBool(strings.ReplaceAll(text, "\n", ""))
		if err != nil {
			continue
		}
		v.Clock24Hour = b
		break
	}

	if err := c.v.ConfigureSettings(&v); err != nil {
		fmt.Println("settings changes have failed: ", err)
	}
}
