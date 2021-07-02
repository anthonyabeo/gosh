package parser

import "testing"

func TestParser(t *testing.T) {
	input := "ls -al a* | grep foo > outfile"
	tests := []Token{
		{value: "ls", typ: IDENTIFIER},
		{value: "-al", typ: OPTION},
		{value: "a*", typ: IDENTIFIER},
		{value: "|", typ: PIPE},
		{value: "grep", typ: IDENTIFIER},
		{value: "foo", typ: IDENTIFIER},
		{value: ">", typ: GREAT},
		{value: "outfile", typ: IDENTIFIER},
	}

	parser := NewParser(input)

	for _, tok := range tests {
		if parser.curToken.typ != tok.typ {
			t.Errorf("Incorrect token. Got=%v, Expected=%v", parser.curToken.typ, IDENTIFIER)
		}

		if parser.curToken.value != tok.value {
			t.Errorf("Wrong token value. Got=%v, Expected=ls", parser.curToken.value)
		}

		parser.NextToken()
	}
}
