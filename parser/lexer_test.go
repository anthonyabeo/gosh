package parser

import (
	"testing"
)

func TestTokens(t *testing.T) {
	input := "& | > >> < >& \n -a grep cat find"
	tests := []Token{
		{value: "&", typ: AMPERSAND},
		{value: "|", typ: PIPE},
		{value: ">", typ: GREAT},
		{value: ">>", typ: GREATGREATER},
		{value: "<", typ: LESS},
		{value: ">&", typ: GREATAMPERSAND},
		{value: "\n", typ: NEWLINE},
		{value: "-a", typ: OPTION},
		{value: "grep", typ: IDENTIFIER},
		{value: "cat", typ: IDENTIFIER},
		{value: "find", typ: IDENTIFIER},
	}

	lexer := NewLexer(input)

	for _, tok := range tests {
		token := lexer.NextToken()

		if token.value != tok.value {
			t.Errorf("Token has wrong value. Got=%s, Expected=%s",
				token.value, tok.value)
		}

		if token.typ != tok.typ {
			t.Errorf("Token has wrong classification. Got=%v, Expected=%v",
				token.typ, tok.typ)
		}
	}
}
