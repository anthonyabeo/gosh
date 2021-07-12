package executor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type CompleteCommand struct {
	Stdin        bytes.Buffer
	Stdout       bytes.Buffer
	Stderr       bytes.Buffer
	Background   bool
	Commands     []*exec.Cmd
	NumAvailCmds int
	NumCmds      int
}

func NewCompleteCommand() *CompleteCommand {
	return &CompleteCommand{
		Background: false,
	}
}

func (cc *CompleteCommand) Execute() {
	var buf bytes.Buffer
	cc.Commands[len(cc.Commands)-1].Stdout = &buf

	for i := cc.NumCmds - 1; i >= 0; i -= 1 {
		cmd := cc.Commands[i]
		if i != 0 {
			err := cmd.Start()
			if err != nil {
				fmt.Fprintf(os.Stderr, "\nError - %v\n", err)
				os.Exit(1)
			}
		} else {
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "\nError - %v\n", err)
				os.Exit(1)
			}
		}
	}

	for i := cc.NumCmds - 1; i > 0; i -= 1 {
		cc.Commands[i].Wait()
	}

	buf.WriteByte('\n')
	io.Copy(os.Stdout, &buf)
}
