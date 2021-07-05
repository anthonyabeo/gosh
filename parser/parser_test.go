package parser

import "testing"

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
