package parser

type Parser struct {
	lex       Lexer
	curToken  Token
	nextToken Token
}

func NewParser(input string) Parser {
	parser := Parser{lex: NewLexer(input)}
	parser.NextToken()
	parser.NextToken()

	return parser
}

func (p *Parser) NextToken() {
	p.curToken = p.nextToken
	p.nextToken = p.lex.NextToken()
}

func (p *Parser) CurToken() Token {
	return p.curToken
}

func (p *Parser) CurTokenIs(tt Token) bool {
	return p.curToken.Typ == tt.Typ
}

func (p *Parser) NextTokenIs(tt Token) bool {
	return p.nextToken.Typ == tt.Typ
}

// func (p *Parser) ParseCommand() CompleteCommand {

// }
