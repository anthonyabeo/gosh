package parser

type TokenType byte

const (
	PIPE TokenType = iota
	AMPERSAND
	GREAT
	GREATGREATER
	IDENTIFIER
	LESS
	GREATAMPERSAND
	NEWLINE
)

type Token struct {
	value              string
	typ                TokenType
	filepath           string
	lineNum, columnNum uint32
}

func NewToken(value string, typ TokenType) Token {
	return Token{
		value:     value,
		typ:       typ,
		filepath:  "",
		lineNum:   0,
		columnNum: 0,
	}
}
