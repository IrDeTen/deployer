package usecase

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/IrDeTen/deployer/models"
	"github.com/IrDeTen/deployer/templates"
)

func CmdMenu() {
	var service models.Service = models.Service{}
	var text string
	sortKeys := make([]string, 0)

	for k := range modules {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)

	clearConsole()
	dVBox.Print("Starting deployment of new service ")
	time.Sleep(3 * time.Second)

	for _, key := range sortKeys {
		if modules[key](&service) {
			clearConsole()
			return
		}
	}
	clearConsole()
loop:
	for {
		sVBox.Print(service.FormString())
		dVBox.Print("Do you want to change new service? (Y/N)")
		fmt.Scan(&text)
		if text == "--exit" {
			clearConsole()
			return
		}
		switch text {
		case "Y", "y":
			clearConsole()
			for {
				t := "Select what you want to change (type '--stop' to stop changes):"
				for key := range modules {
					t = t + "\n  " + key
				}
				dVBox.Print(t)
				fmt.Scan(&text)
				if text == "--stop" {
					clearConsole()
					break
				}
				for key := range modules {
					if strings.Contains(text, key) {
						if modules[key](&service) {
							return
						} else {
							break
						}
					}
				}
			}
		case "N", "n":
			if err := createUnit(&service); err != nil {
				clearConsole()
				errVBox.Print(err.Error())
				continue
			}
			break loop
		}
	}
	if err := runUnit(&service); err != nil {
		clearConsole()
		errVBox.Print(err.Error())
	}
	clearConsole()
	status := exec.Command("systemctl", "status", service.Name)
	stdout, err := status.Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	sVBox.Print(string(stdout))

	dVBox.Print("New service created")
	fmt.Scan(&text)
	if strings.Contains(text, "--exit") {
		clearConsole()
		return
	}

}

func clearConsole() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func createUnit(service *models.Service) error {
	path := "/etc/systemd/system/" + service.Name + ".service"
	t := template.Must(template.New("unit").Parse(templates.Unit))
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, service)
	if err != nil {
		return err
	}
	return nil
}

func runUnit(service *models.Service) error {
	var err error
	reload := exec.Command("systemctl", "daemon-reload")
	err = reload.Run()
	if err != nil {
		return err
	}

	start := exec.Command("systemctl", "start", service.Name)
	err = start.Run()
	if err != nil {
		return err
	}
	return nil
}
