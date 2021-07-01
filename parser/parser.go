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
