package parser

import (
	"os/exec"

	"github.com/anthonyabeo/gosh/executor"
)

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

func (p *Parser) CurTokenIs(tt Token) bool {
	return p.curToken.Typ == tt.Typ
}

func (p *Parser) NextTokenIs(tt Token) bool {
	return p.nextToken.Typ == tt.Typ
}

func (p *Parser) ParseCommand() executor.CompleteCommand {
	cc := executor.CompleteCommand{}

	for p.curToken.Typ != EOF {
		cmd, err := p.parseCmd()
		if err != nil {
			cc.Commands = append(cc.Commands, cmd)
		}

		p.NextToken()
	}

	return cc
}

func (p *Parser) parseCmd() (exec.Cmd, error) {
	cmd := exec.Cmd{Path: p.curToken.Value}
	p.NextToken()

	for p.curToken.Typ == IDENTIFIER || p.curToken.Typ == OPTION {
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.NextToken()
	}

	return cmd, nil
}
