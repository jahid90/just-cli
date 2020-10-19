package parser

import (
	"fmt"

	"github.com/jahid90/just/lib"
	"github.com/jahid90/just/lib/lexer"
)

// Parser The parser
type Parser struct {
	buffer *lexer.TokenBuffer
}

// Returns (position, true), where postition is the index where the token was found in the buffer.
// Returns (-1,false) if not found
// Could probably do a binary search, but there are only a few tokens
func (p *Parser) find(tt lexer.TokenType) (int, bool) {

	var i int
	for {
		token := p.buffer.PeekN(i)

		if token.Type == lexer.EOF {
			return -1, false
		}

		if token.Type == tt {
			if lib.DEBUG {
				fmt.Println("Found " + tt.String() + " at pos: " + fmt.Sprint(i))
			}
			return i, true
		}

		i++
	}
}

// Parse Parses the tokens and generates an AST
func (p *Parser) Parse() *CommandNode {

	if lib.DEBUG {
		fmt.Println("== parse ==")
		p.buffer.Print()
	}

	if p.buffer.HasNext() == false {
		return &CommandNode{
			andCommand: nil,
			orCommand:  nil,
			command:    nil,
		}
	}

	command, p := p.parseCommand()

	if p.buffer.HasNext() {
		fmt.Println("Some tokens were not processed!")
		p.buffer.Print()
	}

	return command
}

func (p *Parser) parseCommand() (*CommandNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseCommand ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	andCommand, p := p.parseAnd()
	if andCommand != nil {
		return &CommandNode{
			andCommand: andCommand,
			orCommand:  nil,
			command:    nil,
		}, p
	}

	orCommand, p := p.parseOr()
	if orCommand != nil {
		return &CommandNode{
			andCommand: nil,
			orCommand:  orCommand,
			command:    nil,
		}, p
	}

	command, p := p.parseCommandExpression()
	if command != nil {
		return &CommandNode{
			andCommand: nil,
			orCommand:  nil,
			command:    command,
		}, p
	}

	return &CommandNode{
		andCommand: nil,
		orCommand:  nil,
		command:    nil,
	}, p
}

func (p *Parser) parseBinaryOperator(tt lexer.TokenType) (*CommandNode, *CommandNode, *Parser) {

	if !p.buffer.HasNext() {
		return nil, nil, p
	}

	pos, ok := p.find(tt)
	if !ok {
		return nil, nil, p
	}

	left, _ := NewParser(p.buffer.TakeBetween(p.buffer.Start(), pos)).parseCommand()
	right, _ := NewParser(p.buffer.TakeBetween(pos+1, p.buffer.End())).parseCommand()
	p.buffer.ForwardToEnd()

	return left, right, p
}

func (p *Parser) parseAnd() (*AndCommandNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseAnd ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	l, r, p := p.parseBinaryOperator(lexer.AND)
	if l == nil && r == nil {
		return nil, p
	}

	return &AndCommandNode{left: l, right: r}, p
}

func (p *Parser) parseOr() (*OrCommandNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseOr ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	l, r, p := p.parseBinaryOperator(lexer.OR)
	if l == nil && r == nil {
		return nil, p
	}

	return &OrCommandNode{left: l, right: r}, p
}

func (p *Parser) parseCommandExpression() (*CommandExpressionNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseCommandExpression ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	env, _p := p.parseEnv()
	if env != nil {
		exec, _pp := _p.parseExec()
		if exec != nil {
			args, _ppp := _pp.parseArgs()

			if args != nil {
				return &CommandExpressionNode{
					env:  env,
					exec: exec,
					args: args,
				}, _ppp
			}

			return &CommandExpressionNode{
				env:  env,
				exec: exec,
				args: nil,
			}, _pp
		}

		panic("Only env variables without a command not expected")
	}

	exec, _p := p.parseExec()
	if p != nil {
		args, _pp := _p.parseArgs()

		if args != nil {
			return &CommandExpressionNode{
				env:  nil,
				exec: exec,
				args: args,
			}, _pp
		}

		return &CommandExpressionNode{
			env:  nil,
			exec: exec,
			args: nil,
		}, _p
	}

	return &CommandExpressionNode{
		env:  nil,
		exec: exec,
		args: nil,
	}, p
}

func (p *Parser) parseEnv() (*EnvNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseEnv ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	multi, p := p.parseMultiEnv()

	return &EnvNode{multi: multi}, p
}

func (p *Parser) parseMultiEnv() (*MultiEnvNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseMultiEnv ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	pos, ok := p.find(lexer.COMMA)
	if !ok {
		keyVal, p := p.parseKeyVal()

		return &MultiEnvNode{keyVal: keyVal, env: nil}, p
	}

	keyVal, _ := NewParser(p.buffer.TakeBetween(p.buffer.Start(), pos)).parseKeyVal()
	env, p := NewParser(p.buffer.TakeBetween(pos+1, p.buffer.End())).parseEnv()

	return &MultiEnvNode{keyVal: keyVal, env: env}, p
}

func (p *Parser) parseKeyVal() (*KeyValNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseKeyVal ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	pos, ok := p.find(lexer.ASSIGN)
	if !ok {
		return nil, p
	}

	key, _ := p.parseIdent()
	val, p := NewParser(p.buffer.TakeBetween(pos+1, p.buffer.End())).parseEnvVal()

	return &KeyValNode{key: key, value: val}, p
}

func (p *Parser) parseEnvVal() (*EnvValNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseEnvVal ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	command, _p := p.parseExpression()
	if command != nil {
		return &EnvValNode{
			command: command,
			value:   nil,
		}, _p
	}

	value, p := p.parseIdent()

	return &EnvValNode{
		command: nil,
		value:   value,
	}, p
}

func (p *Parser) parseExpression() (*CommandNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseExpression ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	pos, ok := p.find(lexer.EXPRS)
	if !ok {
		return nil, p
	}

	end, ok := p.find(lexer.EXPRE)
	if !ok {
		panic("Found expr start, but no end!")
	}

	command, _ := NewParser(p.buffer.TakeBetween(pos+1, end)).parseCommand()
	p = NewParser(p.buffer.TakeBetween(end+1, p.buffer.End()))

	return command, p
}

func (p *Parser) parseExec() (*ExecNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseExec ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	value, _ := p.parseIdent()

	return &ExecNode{value: value}, p
}

func (p *Parser) parseArgs() (*ArgNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseArgs ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	multi, p := p.parseMultiArg()

	return &ArgNode{multi: multi}, p
}

func (p *Parser) parseMultiArg() (*MultiArgNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseMultiArg ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	// There can be an arg before the expression
	pos, ok := p.find(lexer.EXPRS)
	if ok {
		if pos != p.buffer.Start() {
			// there's an arg!
			value, _p := p.parseIdent()

			if value != nil {

				args, _pp := _p.parseArgs()

				if args != nil {
					return &MultiArgNode{
						command: nil,
						arg:     args,
						value:   value,
					}, _pp
				}

				return &MultiArgNode{
					command: nil,
					arg:     nil,
					value:   value,
				}, _p
			}

			panic("There was something before the expression, but not an arg!")

		}
	}

	command, _p := p.parseExpression()
	if command != nil {
		args, _pp := _p.parseArgs()

		if args != nil {
			return &MultiArgNode{
				command: command,
				arg:     args,
				value:   nil,
			}, _pp
		}

		return &MultiArgNode{
			command: command,
			arg:     nil,
			value:   nil,
		}, _p
	}

	value, _p := p.parseIdent()
	if value != nil {
		args, _pp := _p.parseArgs()

		if args != nil {
			return &MultiArgNode{
				command: nil,
				arg:     args,
				value:   value,
			}, _pp
		}

		return &MultiArgNode{
			command: nil,
			arg:     nil,
			value:   value,
		}, _p
	}

	return &MultiArgNode{
		command: nil,
		arg:     nil,
		value:   nil,
	}, p
}

func (p *Parser) parseIdent() (*IdentNode, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseIdent ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return nil, p
	}

	value, _ := p.parseValue()

	return &IdentNode{value: value}, p
}

func (p *Parser) parseValue() (string, *Parser) {

	if lib.DEBUG {
		fmt.Println("== parseValue ==")
		p.buffer.Print()
	}

	if !p.buffer.HasNext() {
		return "nil", p
	}

	value := p.buffer.Next().Value

	return value, p
}

// NewParser Creates an instance of a parser
func NewParser(tb *lexer.TokenBuffer) *Parser {
	return &Parser{
		buffer: tb,
	}
}
