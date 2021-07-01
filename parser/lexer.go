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
	case '|':
		token.typ = PIPE
		token.value = "|"
	case '>':
		if lex.nextCharIs('>') {
			token.typ = GREATGREATER
			token.value = ">>"

			lex.readChar()
		} else if lex.nextCharIs('&') {
			token.typ = GREATAMPERSAND
			token.value = ">&"

			lex.readChar()
		} else {
			token.typ = GREAT
			token.value = ">"
		}
	case '<':
		token.typ = LESS
		token.value = "<"
	case '\n':
		token.typ = NEWLINE
		token.value = "\n"
	case '-':
		token.typ = OPTION
		token.value = lex.readIdentifier()
	default:
		return Token{}
	}

	lex.readChar()

	return token
}

func (lex *Lexer) nextCharIs(chr byte) bool {
	if lex.nextCharPos >= len(lex.input) {
		return false
	}

	return lex.input[lex.nextCharPos] == chr
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
	for lex.curChar == ' ' || lex.curChar == '\t' || lex.curChar == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) readIdentifier() string {
	pos := lex.curCharPos
	for lex.curCharIsLetter() {
		lex.readChar()
	}

	return lex.input[pos:lex.curCharPos]
}

func (lex *Lexer) curCharIsLetter() bool {
	return ('a' <= lex.curChar && lex.curChar <= 'z') ||
		('A' <= lex.curChar && lex.curChar <= 'Z') ||
		lex.curChar == '_' ||
		lex.curChar == '-'
}
