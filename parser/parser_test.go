package parser

import (
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

	if cmd.Path != "ls" {
		t.Errorf("Wrong command path. Got=%v, Expected=ls", cmd.Path)
	}

	if len(cmd.Args) != 2 {
		t.Errorf("Wrong number of arguments and options. Got=%v, Expected=2", len(cmd.Args))
	}

	if cmd.Args[0] != "-al" {
		t.Errorf("Wrong first option. Got=%v, Expected=-al", cmd.Args[0])
	}

	if cmd.Args[1] != "foo" {
		t.Errorf("Wrong second option. Got=%v, Expected=foo", cmd.Args[1])
	}
}
