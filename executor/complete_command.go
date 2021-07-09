package executor

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type CompleteCommand struct {
	Stdin        io.Reader
	Stdout       io.Writer
	Stderr       io.Writer
	Background   bool
	Commands     []*exec.Cmd
	NumAvailCmds uint
	NumCmds      uint
}

func NewCompleteCommand() *CompleteCommand {
	return &CompleteCommand{
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		Background: false,
	}
}

func (cc *CompleteCommand) Execute() {
	for _, cmd := range cc.Commands {
		cmdOut, err := cmd.Output()

		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError - %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s\n", string(cmdOut))
	}
}
