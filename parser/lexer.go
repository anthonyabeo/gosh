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
		token.Typ = AMPERSAND
		token.Value = "&"
	case '|':
		token.Typ = PIPE
		token.Value = "|"
	case '>':
		if lex.nextCharIs('>') {
			token.Typ = GREATGREATER
			token.Value = ">>"

			lex.readChar()
		} else if lex.nextCharIs('&') {
			token.Typ = GREATAMPERSAND
			token.Value = ">&"

			lex.readChar()
		} else {
			token.Typ = GREAT
			token.Value = ">"
		}
	case '<':
		token.Typ = LESS
		token.Value = "<"
	case '\n':
		token.Typ = NEWLINE
		token.Value = "\n"
	case '-':
		token.Typ = OPTION
		token.Value = lex.readIdentifier()
	default:
		if lex.curCharIsLetter() {
			token.Typ = IDENTIFIER
			token.Value = lex.readIdentifier()
		} else {
			return Token{Typ: EOF}
		}
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
		lex.curChar == '-' ||
		lex.curChar == '*' ||
		lex.curChar == '.'
}
