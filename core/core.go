package core

import (
	"os"
	"os/exec"
)

type OS uint8

const (
	WINDOWS OS = iota
	LINUX
	MAC
)

type CommandBuilder interface {
	BuildCommand() string
	SetDefaults()
	GetOS() OS
	SetURL(string)
	Clone() CommandBuilder
}

// BuildCommand takes in a CommandBuilder (either VideoOptions, AudioOptions, or PlaylistOptions)
func BuildCommand(opts CommandBuilder) *exec.Cmd {
	cmdStr := opts.BuildCommand()

	shell, flag := "bash", "-c" //both for linux and mac. change this if we ever wanna add another OS (why would we ever do that xdxd?)
	if opts.GetOS() == WINDOWS {
		shell, flag = "cmd", "\\C"
	}
	cmd := exec.Command(shell, flag, cmdStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}
