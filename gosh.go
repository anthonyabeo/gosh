package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"github.com/anthonyabeo/gosh/parser"
)

func main() {
	// TODO - Implement Shell builtin commands
	// TODO - Wildcards and Shell expansions
	// TODO - POSIX Compliant
	var username, hostname, cwd string
	var u user.User

	if u, err := user.Current(); err != nil {
		username = "vagrant"
	} else {
		username = u.Username
	}

	if hname, err := os.Hostname(); err != nil {
		hostname = "localhos"
	} else {
		hostname = hname
	}

	if pwd, err := os.Getwd(); err != nil {
		cwd = u.HomeDir
	} else {
		cwd = pwd
	}

	prompt := fmt.Sprintf("[%s@%s:%s]$ ", username, hostname, cwd)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)

		scanner.Scan()
		text := scanner.Text()

		if len(text) > 0 {
			p := parser.NewParser(text)
			cc := p.ParseCommand()

			cc.Execute()
		} else {
			fmt.Fprint(os.Stdout, "\n")
		}
	}
}
