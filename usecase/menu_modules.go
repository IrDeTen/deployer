package usecase

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/IrDeTen/deployer/models"
)

var modules map[string]func(*models.Service) bool = map[string]func(*models.Service) bool{
	"01 name":              setName,
	"02 description":       serDescription,
	"03 exec path":         setExecPath,
	"04 working directory": setWorkDirectory,
	"05 type":              setType,
	"06 after services":    setAfterServices,
	"07 required services": setReqServices,
	"08 restart":           setRestart,
	"09 user":              setUser,
	"10 group":             setGroup,
}

func setName(service *models.Service) bool {
	var text string
	clearConsole()
	for {
		dVBox.Print("Enter new service name")
		fmt.Scan(&text)
		if len(text) == 0 {
			clearConsole()
			errVBox.Print("New service name is blank")
			continue
		}
		if strings.Contains(text, "--exit") {
			return true
		}
		service.Name = text
		if err := service.CheckName(); err != nil {
			clearConsole()
			errVBox.Print(err.Error())
			continue
		}
		break
	}
	return false
}

func serDescription(service *models.Service) bool {
	var text string
	clearConsole()
	sVBox.Print(service.FormString())
	dVBox.Print("Enter new service description")
	fmt.Scan(&text)

	if strings.Contains(text, "--exit") {
		return true
	}
	service.Description = text
	return false
}

func setExecPath(service *models.Service) bool {
	var text string
	clearConsole()
	for {
		sVBox.Print(service.FormString())
		dVBox.Print("Enter path to new service executable")
		fmt.Scan(&text)
		if len(text) == 0 {
			clearConsole()
			errVBox.Print("Path is blank")
			continue
		}
		if strings.Contains(text, "--exit") {
			return true
		}
		service.Path = text
		if err := service.CheckPath("exec"); err != nil {
			clearConsole()
			errVBox.Print(err.Error())
			continue
		}
		break
	}
	return false
}

func setWorkDirectory(service *models.Service) bool {
	var text string
	split := strings.SplitAfter(service.Path, "/")
	if len(split) == 1 {
		service.WorkingDirectory = split[0]
	} else {
		for i := 0; i < len(split)-1; i++ {

			service.WorkingDirectory = service.WorkingDirectory + split[i]
		}
	}

	clearConsole()
	sVBox.Print(service.FormString())
	dVBox.Print(fmt.Sprintf("Working directory by default is\n%s\n\nDo you want to change it? (Y/N)",
		service.WorkingDirectory))
	fmt.Scan(&text)
	if strings.Contains(text, "--exit") {
		return true
	}
	if strings.Contains(text, "Y") ||
		strings.Contains(text, "y") {
		clearConsole()
		for {
			dVBox.Print("Enter path to working directory for service")
			fmt.Scan(&text)
			if strings.Contains(text, "--exit") {
				return true
			}
			if len(text) == 0 {
				clearConsole()
				errVBox.Print("New path is empty")
				continue
			}
			service.WorkingDirectory = text
			if err := service.CheckPath("work"); err != nil {
				clearConsole()
				errVBox.Print(err.Error())
				continue
			}
			break
		}
	}
	return false
}

func setType(service *models.Service) bool {
	var text string
	clearConsole()
	for {
		sVBox.Print(service.FormString())
		dVBox.Print("Select type of new service:\n  simple\n  exec\n  forking\n  oneshot\n  dbus\n  notify\n  idle")
		fmt.Scan(&text)

		switch text {
		case "--exit":
			return true
		case "simple", "exec", "forking",
			"oneshot", "dbus", "notify", "idle":
			service.Type = text
		default:
			clearConsole()
			errVBox.Print("Wrong service type")
			continue
		}
		break
	}
	return false
}

func setAfterServices(service *models.Service) bool {
	var text string
	clearConsole()
	sVBox.Print(service.FormString())
	dVBox.Print("Is service need other services start previosly? (Y/N)")
	fmt.Scan(&text)
	if strings.Contains(text, "--exit") {
		return true
	}
	if strings.Contains(text, "Y") ||
		strings.Contains(text, "y") {
		clearConsole()
		for {
			sVBox.Print(service.FormString())
			dVBox.Print("Enter service name (type '--stop' to finish adding services)")
			fmt.Scan(&text)
			if strings.Contains(text, "--exit") {
				return true
			}
			if strings.Contains(text, "--stop") {
				break
			}
			if len(text) == 0 {
				clearConsole()
				errVBox.Print("Name is blank")
				continue
			}
			if err := service.CheckAndAppendService("after", text); err != nil {
				clearConsole()
				errVBox.Print(err.Error())
				continue
			}
			clearConsole()
		}
	}
	return false
}

func setReqServices(service *models.Service) bool {
	var text string
	clearConsole()
	sVBox.Print(service.FormString())
	dVBox.Print("Is service requiers other services? (Y/N)")
	fmt.Scan(&text)
	if strings.Contains(text, "--exit") {
		return true
	}
	if strings.Contains(text, "Y") ||
		strings.Contains(text, "y") {
		clearConsole()
		for {
			sVBox.Print(service.FormString())
			dVBox.Print("Enter service name (type '--stop' to finish adding services)")
			fmt.Scan(&text)
			if strings.Contains(text, "--exit") {
				return true
			}
			if strings.Contains(text, "--stop") {
				break
			}
			if len(text) == 0 {
				clearConsole()
				errVBox.Print("Name is blank")
				continue
			}
			if err := service.CheckAndAppendService("requires", text); err != nil {
				clearConsole()
				errVBox.Print(err.Error())
				continue
			}
			clearConsole()
		}
	}
	return false

}

func setRestart(service *models.Service) bool {
	var text string
	clearConsole()
	for {
		sVBox.Print(service.FormString())
		dVBox.Print("Select service restart option:\n  no\n  on-success\n  on-failure\n  on-abnormal\n  on-watchdog\n  on-abort\n  always")
		fmt.Scan(&text)

		switch text {
		case "--exit":
			return true
		case "no", "on-success", "on-failure",
			"on-abnormal", "on-watchdog", "on-abort", "always":
			service.Restart = text
		default:
			clearConsole()
			errVBox.Print("Wrong restart option")
			continue
		}
		break
	}
	return false
}

func setUser(service *models.Service) bool {
	var text string
	clearConsole()
	getUsers := exec.Command("cut", "-d:", "-f1", "/etc/passwd")
	stdOut, err := getUsers.Output()
	if err != nil {
		clearConsole()
		errVBox.Print(err.Error())
	}
	for {
		dVBox.Print(service.FormString())
		sVBox.Print("Select user for running srevice:\n" + string(stdOut))
		fmt.Scan(&text)
		if text == "--exit" {
			clearConsole()
			return true
		}
		if len(text) == 0 {
			clearConsole()
			errVBox.Print("User name is blank")
			continue
		}
		service.User = text
		break
	}
	return false
}

func setGroup(service *models.Service) bool {
	var text string
	clearConsole()
	getGroups := exec.Command("groups", service.User)
	stdOut, err := getGroups.Output()
	if err != nil {
		clearConsole()
		errVBox.Print(err.Error())
	}
	groups := strings.Split(
		strings.Split(string(stdOut), " : ")[1],
		"\n",
	)
	for {
		dVBox.Print(service.FormString())
		sVBox.Print("Select user for running srevice:\n" + strings.Join(groups, "\n"))
		fmt.Scan(&text)
		if text == "--exit" {
			clearConsole()
			return true
		}
		if len(text) == 0 {
			clearConsole()
			errVBox.Print("User name is blank")
			continue
		}
		service.Group = text
		break
	}
	return false
}
