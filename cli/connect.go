package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func (c *conf) vectorConnect(args []string) {
	if c.v.Connected() {
		fmt.Println("bluetooth connectivity has already been established")
		return
	}

	//nolint
	if len(args) != 2 {
		fmt.Println("invalid argument.  Usage is \n connect ID")
		return
	}

	i, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid argument.  Usage is \n connect ID")
	}

	if err := c.v.Connect(i); err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var pin string

	fmt.Println("Once the pin comes up on your vector screen, enter it here")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		pin = strings.ReplaceAll(text, "\n", "")
		break
	}

	if err := c.v.SendPin(pin); err != nil {
		log.Println("unable to set pin: ", err)
		return
	}

}
