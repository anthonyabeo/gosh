package parser

import (
	"testing"
)

func TestTokens(t *testing.T) {
	input := "& | > >> < &> \n -a grep cat find"
	tests := []Token{
		{Value: "&", Typ: AMPERSAND},
		{Value: "|", Typ: PIPE},
		{Value: ">", Typ: GREAT},
		{Value: ">>", Typ: GREATGREATER},
		{Value: "<", Typ: LESS},
		{Value: "&>", Typ: AMPERSANDGREAT},
		{Value: "\n", Typ: NEWLINE},
		{Value: "-a", Typ: OPTION},
		{Value: "grep", Typ: IDENTIFIER},
		{Value: "cat", Typ: IDENTIFIER},
		{Value: "find", Typ: IDENTIFIER},
	}

	lexer := NewLexer(input)

	for _, tok := range tests {
		token := lexer.NextToken()

		if token.Value != tok.Value {
			t.Errorf("Token has wrong Value. Got=%s, Expected=%s",
				token.Value, tok.Value)
		}

		if token.Typ != tok.Typ {
			t.Errorf("Token has wrong classification. Got=%v, Expected=%v",
				token.Typ, tok.Typ)
		}
	}
}
