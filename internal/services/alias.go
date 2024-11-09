package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Alias struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Path    string `json:"path"`
}

func (a *Alias) Run(args []string, projectPath string) error {
	path := projectPath + "/" + a.Path
	args = append([]string{a.Command}, args...)

	cmd := exec.Command(Cfg.Shell, "-c", strings.Join(args, " "))
	cmd.Dir = path

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and check for errors
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}
