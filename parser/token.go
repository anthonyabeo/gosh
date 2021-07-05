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
	OPTION
	ILLEGAL
	EOF
)

type Token struct {
	Value              string
	Typ                TokenType
	Filepath           string
	LineNum, ColumnNum uint32
}

func NewToken(value string, typ TokenType) Token {
	return Token{
		Value:     value,
		Typ:       typ,
		Filepath:  "",
		LineNum:   0,
		ColumnNum: 0,
	}
}
