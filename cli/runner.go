package cli

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunCommand executes the given shell command string.
func RunCommand(cmdStr string) error {
	// Split command string into command and args
	cmdArgs := splitCommand(cmdStr)
	if len(cmdArgs) == 0 {
		return fmt.Errorf("empty command")
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil
	cmd.Stdout = exec.Command("tee", "/dev/stderr").Stdout
	cmd.Stderr = exec.Command("tee", "/dev/stderr").Stderr

	// Run command and wait
	return cmd.Run()
}

// splitCommand splits a command string into command and args respecting quoted strings.
func splitCommand(cmd string) []string {
	// Simple split by spaces - for production use a shellwords parser if needed.
	return strings.Fields(cmd)
}
