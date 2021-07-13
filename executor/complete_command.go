package executor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type CompleteCommand struct {
	Stdin                         bytes.Buffer
	Stdout                        bytes.Buffer
	Stderr                        bytes.Buffer
	Background                    bool
	Commands                      []*exec.Cmd
	NumAvailCmds                  int
	NumCmds                       int
	StdoutFilename, StdinFilename string
}

func NewCompleteCommand() *CompleteCommand {
	return &CompleteCommand{
		Background: false,
	}
}

func (cc *CompleteCommand) Execute() {
	var buf, inputBuf bytes.Buffer

	if len(cc.StdinFilename) > 0 {
		fileContent, readErr := ioutil.ReadFile(cc.StdinFilename)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "\nError - %v\n", readErr)
			os.Exit(1)
		}

		numBytes, err := inputBuf.Write(fileContent)
		if err != nil || numBytes != len(fileContent) {
			log.Fatal(err)
		} else {
			cc.Commands[0].Stdin = &inputBuf
		}
	}

	cc.Commands[len(cc.Commands)-1].Stdout = &buf

	for i := cc.NumCmds - 1; i >= 0; i -= 1 {
		cmd := cc.Commands[i]
		var err error

		if i > 0 {
			err = cmd.Start()
		} else {
			err = cmd.Run()
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError - %v\n", err)
			os.Exit(1)
		}
	}

	for i := cc.NumCmds - 1; i > 0; i -= 1 {
		cc.Commands[i].Wait()
	}

	buf.WriteByte('\n')

	if len(cc.StdoutFilename) > 0 {
		outfile, err := os.Create(cc.StdoutFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError - %v\n", err)
			os.Exit(1)
		}

		writer := bufio.NewWriter(outfile)
		writer.WriteString(buf.String())
		writer.Flush()

		fmt.Fprint(os.Stdout, "\n")
	} else {
		io.Copy(os.Stdout, &buf)
	}
}
