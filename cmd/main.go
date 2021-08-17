package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menu struct {
	reader *bufio.Reader
}

func main() {
	var cmdMenu menu

	cmdMenu.reader = bufio.NewReader(os.Stdin)

	for {
		cmdMenu.CmdMenu()
	}

}

func (m *menu) CmdMenu() {
	var service models.Service

	fmt.Print("Service Deployer is running \nEnter the name of new service\n-> ")
	text, _ := m.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

}
