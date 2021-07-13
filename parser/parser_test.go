package parser

import (
	"os/exec"
	"testing"
)

func TestParser(t *testing.T) {
	input := "ls -al a* | grep foo > outfile"
	tests := []Token{
		{Value: "ls", Typ: IDENTIFIER},
		{Value: "-al", Typ: OPTION},
		{Value: "a*", Typ: IDENTIFIER},
		{Value: "|", Typ: PIPE},
		{Value: "grep", Typ: IDENTIFIER},
		{Value: "foo", Typ: IDENTIFIER},
		{Value: ">", Typ: GREAT},
		{Value: "outfile", Typ: IDENTIFIER},
	}

	parser := NewParser(input)

	for _, tok := range tests {
		if parser.curToken.Typ != tok.Typ {
			t.Errorf("Incorrect token. Got=%v, Expected=%v", parser.curToken.Typ, IDENTIFIER)
		}

		if parser.curToken.Value != tok.Value {
			t.Errorf("Wrong token Value. Got=%v, Expected=ls", parser.curToken.Value)
		}

		parser.NextToken()
	}
}

func TestParseCmd(t *testing.T) {
	input := "ls -al foo"
	p := NewParser(input)
	cmd, err := p.parseCmd()

	if err != nil {
		t.Errorf("Cmd parse failed. Expected err to be nil.")
	}

	lsBin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		t.Error("Lookup for ls binary failed")
	}

	if cmd.Path != lsBin {
		t.Errorf("Wrong command path. Got=%v, Expected=%v", cmd.Path, lsBin)
	}

	if len(cmd.Args) != 3 {
		t.Errorf("Wrong number of arguments and options. Got=%v, Expected=3", len(cmd.Args))
	}

	if cmd.Args[1] != "-al" {
		t.Errorf("Wrong first option. Got=%v, Expected=-al", cmd.Args[0])
	}

	if cmd.Args[2] != "foo" {
		t.Errorf("Wrong second option. Got=%v, Expected=foo", cmd.Args[1])
	}
}

func TestParseCommand(t *testing.T) {
	input := "ls -al foo"
	p := NewParser(input)
	cc := p.ParseCommand()

	if len(cc.Commands) != 1 {
		t.Errorf("Wrong number of commands. Got=%v, Expected=1",
			len(cc.Commands))
	}

	lsBin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		t.Error("Lookup for ls binary failed")
	}

	if cc.Commands[0].Path != lsBin {
		t.Errorf("Wrong path of first command. Got=%v, Expected=%v",
			cc.Commands[0].Path, lsBin)
	}

}

func TestPipedCommand(t *testing.T) {
	input := "ls -al | grep git"

	p := NewParser(input)
	cc := p.ParseCommand()

	// /bin/ls Subcommand Tests
	lsBin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		t.Error("Lookup for ls binary failed")
	}

	if cc.Commands[0].Path != lsBin {
		t.Errorf("Wrong path for first subcommand. Got=%v, Expected=%v",
			cc.Commands[0].Path, lsBin)
	}

	if len(cc.Commands[0].Args) != 2 {
		t.Errorf("Wrong number of args. Got=%v, Expected=2",
			len(cc.Commands[0].Args))
	}

	// /usr/bin/grep Subcommand Tests
	grepBin, grepErr := exec.LookPath("grep")
	if grepErr != nil {
		t.Error("Lookup for grep binary failed")
	}

	if cc.Commands[1].Path != grepBin {
		t.Errorf("Wrong path for first subcommand. Got=%v, Expected=%v",
			cc.Commands[1].Path, grepBin)
	}

	if len(cc.Commands[1].Args) != 2 {
		t.Errorf("Wrong number of args. Got=%v, Expected=2",
			len(cc.Commands[1].Args))
	}

	// Full Command Tests
	if cc.NumCmds != 2 {
		t.Errorf("wrong number of subcommand. Got=%v, Expected=2", cc.NumCmds)
	}

	if cc.Background {
		t.Errorf("Background execution not specfied. Got=%v, Expected=false",
			cc.Background)
	}
}

func TestOutputRedirection(t *testing.T) {
	input := "ls -al | grep git | wc > outfile.txt"
	p := NewParser(input)
	cc := p.ParseCommand()

	// /bin/ls Subcommand Tests
	lsBin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		t.Error("Lookup for ls binary failed")
	}

	if cc.Commands[0].Path != lsBin {
		t.Errorf("Wrong path for first subcommand. Got=%v, Expected=%v",
			cc.Commands[0].Path, lsBin)
	}

	if len(cc.Commands[0].Args) != 2 {
		t.Errorf("Wrong number of args. Got=%v, Expected=2",
			len(cc.Commands[0].Args))
	}

	// /usr/bin/grep Subcommand Tests
	grepBin, grepErr := exec.LookPath("grep")
	if grepErr != nil {
		t.Error("Lookup for grep binary failed")
	}

	if cc.Commands[1].Path != grepBin {
		t.Errorf("Wrong path for first subcommand. Got=%v, Expected=%v",
			cc.Commands[1].Path, grepBin)
	}

	if len(cc.Commands[1].Args) != 2 {
		t.Errorf("Wrong number of args. Got=%v, Expected=2",
			len(cc.Commands[1].Args))
	}

	wcBin, wcErr := exec.LookPath("wc")
	if wcErr != nil {
		t.Error("Lookup for wc binary failed")
	}

	if cc.Commands[2].Path != wcBin {
		t.Errorf("Wrong path for first subcommand. Got=%v, Expected=%v",
			cc.Commands[2].Path, wcBin)
	}

	if len(cc.Commands[2].Args) != 1 {
		t.Errorf("Wrong number of args. Got=%v, Expected=1",
			len(cc.Commands[2].Args))
	}

	// Full Command Tests
	if cc.NumCmds != 3 {
		t.Errorf("wrong number of subcommand. Got=%v, Expected=3", cc.NumCmds)
	}

	if cc.Background {
		t.Errorf("Background execution not specfied. Got=%v, Expected=false",
			cc.Background)
	}

	if cc.StdoutFilename != "outfile.txt" {
		t.Errorf("Wrong output file name. Got=%v, Expected=outfile.txt",
			cc.StdoutFilename)
	}
}
