package parser

import (
	"log"
	"os/exec"

	"github.com/anthonyabeo/gosh/executor"
)

type Parser struct {
	lex      Lexer
	CurToken Token
	NxtToken Token
}

func NewParser(input string) Parser {
	parser := Parser{lex: NewLexer(input)}
	parser.NextToken()
	parser.NextToken()

	return parser
}

func (p *Parser) NextToken() {
	p.CurToken = p.NxtToken
	p.NxtToken = p.lex.NextToken()
}

func (p *Parser) CurTokenTypeIs(tt TokenType) bool {
	return p.CurToken.Typ == tt
}

func (p *Parser) NextTokenTypeIs(tt TokenType) bool {
	return p.NxtToken.Typ == tt
}

func (p *Parser) ParseCommand() *executor.CompleteCommand {
	cc := executor.NewCompleteCommand()
	pipe := false

	for !p.CurTokenTypeIs(EOF) {
		if p.CurTokenTypeIs(PIPE) {
			pipe = true
			p.NextToken()
		}

		if p.CurTokenTypeIs(GREAT) {
			p.NextToken()
			cc.StdoutFilename = p.CurToken.Value
			p.NextToken()

			continue
		}

		if p.CurTokenTypeIs(GREATGREATER) {
			p.NextToken()
			cc.StdoutFilename = p.CurToken.Value
			cc.AppendOutput = true
			p.NextToken()

			continue
		}

		if p.CurTokenTypeIs(LESS) {
			p.NextToken()
			cc.StdinFilename = p.CurToken.Value
			p.NextToken()

			continue
		}

		if p.CurTokenTypeIs(AMPERSANDGREAT) {
			p.NextToken()
			cc.StdoutFilename = p.CurToken.Value
			cc.MergeOutErr = true
			p.NextToken()

			continue
		}

		if p.CurTokenTypeIs(AMPERSANDGREATGREAT) {
			p.NextToken()
			cc.StdoutFilename = p.CurToken.Value
			cc.MergeOutErr = true
			cc.AppendOutput = true
			p.NextToken()

			continue
		}

		cmd, err := p.parseCmd()
		if err == nil {
			if pipe {
				prevCmd := cc.Commands[len(cc.Commands)-1]
				if prevCmdStdout, err := prevCmd.StdoutPipe(); err != nil {
					log.Fatal(err)
				} else {
					cmd.Stdin = prevCmdStdout
					pipe = false
				}
			}

			cc.Commands = append(cc.Commands, cmd)
			cc.NumCmds += 1
		}

		p.NextToken()
	}

	return cc
}

func (p *Parser) parseCmd() (*exec.Cmd, error) {
	// TODO Implement error cases.
	var args []string
	path := p.CurToken.Value

	for p.NextTokenTypeIs(IDENTIFIER) || p.NextTokenTypeIs(OPTION) {
		p.NextToken()

		args = append(args, p.CurToken.Value)
	}
	cmd := exec.Command(path, args...)

	return cmd, nil
}
