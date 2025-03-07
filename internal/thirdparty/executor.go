// internal/thirdparty/executor.go
package thirdparty

import (
	"docmap-client-proxy-go/internal/logger"
	"os/exec"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Error executing %s: %v", command, err)
		return "", err
	}
	logger.Info("Executed command: %s %v", command, args)
	return string(output), nil
}
