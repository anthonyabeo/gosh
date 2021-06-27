package parser

type Lexer struct {
	lineNum, columnNum      uint32
	curChar                 byte
	input                   string
	curCharPos, nextCharPos int
}

func NewLexer(input string) Lexer {
	lexer := Lexer{
		input:       input,
		lineNum:     0,
		columnNum:   0,
		curChar:     0,
		curCharPos:  0,
		nextCharPos: 0,
	}

	lexer.readChar()

	return lexer
}

func (lex *Lexer) NextToken() Token {
	var token Token
	lex.skipWhiteSpace()

	switch lex.curChar {
	case '&':
		token.typ = AMPERSAND
		token.value = "&"
	default:
		return Token{}
	}

	return token
}

func (lex *Lexer) readChar() {
	if lex.nextCharPos >= len(lex.input) {
		lex.curChar = 0
	} else {
		lex.curChar = lex.input[lex.nextCharPos]
	}

	lex.curCharPos = lex.nextCharPos
	lex.nextCharPos += 1
}

func (lex *Lexer) skipWhiteSpace() {
	for lex.curChar == ' ' || lex.curChar == '\t' || lex.curChar == '\n' || lex.curChar == '\r' {
		lex.readChar()
	}
}