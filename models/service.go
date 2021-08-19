package models

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/IrDeTen/deployer/templates"
)

type Service struct {
	Name             string
	Description      string
	Path             string
	Type             string
	WorkingDirectory string
	After            []string
	Requires         []string
	Restart          string
	User             string
	Group            string
}

func (s *Service) CheckName() error {
	var err error
	if len(s.Name) == 0 {
		return errors.New("error: \nService name is blank")
	}

	path := fmt.Sprintf("/etc/systemd/system/%s.service", s.Name)
	if _, err = os.Stat(path); err == nil {
		return fmt.Errorf("error: \nService with this name alredy exist/ Path: %s", path)
	}

	path = fmt.Sprintf("/lib/systemd/system/%s.service", s.Name)
	if _, err = os.Stat(path); err == nil {
		s.Name = ""
		return fmt.Errorf("error: \nService with this name alredy exist:\n%s", path)
	}
	return nil
}

func (s *Service) CheckPath(key string) error {
	var err error
	switch key {
	case "exec":
		if len(s.Path) == 0 {
			return errors.New("error: \nPath to executable is blank")
		}
		if _, err = os.Stat(s.Path); err != nil {
			s.Path = ""
			return fmt.Errorf("error: \nFailed to find executable in %s", s.Path)
		}
	case "work":
		if len(s.WorkingDirectory) == 0 {
			return errors.New("error: \nPath to working directory is blank")
		}
		if _, err = os.Stat(s.WorkingDirectory); err != nil {
			s.WorkingDirectory = ""
			return fmt.Errorf("error: \nFailed to find directory in %s", s.WorkingDirectory)
		}
	}
	return nil
}

func (s *Service) CheckType() error {
	if len(s.Type) == 0 {
		return errors.New("error: \nService type is blank")
	}
	return nil
}

func (s *Service) CheckAndAppendService(key, serviceName string) error {
	var err error
	etcPath := "/etc/systemd/system/"
	libPath := "/lib/systemd/system/"
	if _, err = os.Stat(etcPath + serviceName); err != nil {
		if _, err = os.Stat(libPath + serviceName); err != nil {
			return fmt.Errorf("error: \nService with this name doesn't found in\n%s\n%s", etcPath, libPath)
		}
	}
	switch key {
	case "after":
		s.After = append(s.After, serviceName)
	case "requires":
		s.Requires = append(s.Requires, serviceName)
	}

	return nil
}

func (s *Service) FormString() string {
	args := []interface{}{
		s.Name,
		s.Description,
		s.Path,
		s.WorkingDirectory,
		s.Type,
		strings.Join(s.After, ", "),
		strings.Join(s.Requires, ", "),
		s.Restart,
		s.User,
		s.Group,
	}
	return fmt.Sprintf(strings.Join(templates.ServiceTmpl, "\n"), args...)
}
