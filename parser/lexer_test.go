package parser

import (
	"testing"
)

func TestTokens(t *testing.T) {
	input := "&"
	tests := []Token{
		{value: "&", typ: AMPERSAND},
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
