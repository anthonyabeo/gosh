package executor

import (
	"io"
	"os/exec"
)

type CompleteCommand struct {
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	Background bool
	Commands   []exec.Cmd
}
