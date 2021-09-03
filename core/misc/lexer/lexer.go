package lexer

import (
	"bufio"
	"io"
	"unicode"
)

/* See lib/parser/grammar.go for the grammar */

// Lexer Represents a lexer
type Lexer struct {
	pos    int
	reader *bufio.Reader
}

// Run Runs the lexer and returns a TokenBuffer
func (l *Lexer) Run() *TokenBuffer {
	var tokens []*Token

	for {
		token := l.Lex()

		tokens = append(tokens, &token)

		if token.Type == EOF {
			break
		}
	}

	return NewTokenBuffer(tokens)
}

// Lex Scans the input for the next token.
// It returns the position and type of the token and the literal value
func (l *Lexer) Lex() Token {
	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return Token{Type: EOF, Position: l.pos, Value: "EOF"}
			}

			panic(err)
		}

		l.pos++

		switch r {
		case '=':
			return Token{Type: ASSIGN, Position: l.pos, Value: "="}

		case ',':
			return Token{Type: COMMA, Position: l.pos, Value: ","}

		case '&':
			startPos := l.pos
			l.unread()
			token, ok := l.lexAnd()
			if !ok {
				return Token{Type: IDENT, Position: startPos, Value: token}
			}
			return Token{Type: AND, Position: startPos, Value: token}

		case '|':
			startPos := l.pos
			l.unread()
			token, ok := l.lexOr()
			if !ok {
				return Token{Type: IDENT, Position: startPos, Value: token}
			}
			return Token{Type: OR, Position: startPos, Value: token}

		case '$':
			startPos := l.pos
			l.unread()
			token, ok := l.lexExprStart()
			if !ok {
				return Token{Type: IDENT, Position: startPos, Value: token}
			}
			return Token{Type: EXPRS, Position: startPos, Value: token}

		case ')':
			return Token{Type: EXPRE, Position: l.pos, Value: ")"}

		default:
			if unicode.IsSpace(r) {
				continue
			} else {
				startPos := l.pos
				l.unread()
				token, ok := l.lexIdent()
				if !ok {
					return Token{Type: UNKNOWN, Position: startPos, Value: ""}
				}
				return Token{Type: IDENT, Position: startPos, Value: token}
			}

		}
	}
}

func (l *Lexer) read() (rune, error) {
	r, _, err := l.reader.ReadRune()

	return r, err
}

func (l *Lexer) unread() error {
	err := l.reader.UnreadRune()
	if err != nil {
		return err
	}

	l.pos--

	return nil
}

func (l *Lexer) lexAnd() (string, bool) {
	return l.lexConsecutive('&')
}

func (l *Lexer) lexOr() (string, bool) {
	return l.lexConsecutive('|')
}

func (l *Lexer) lexConsecutive(r rune) (string, bool) {
	_, err := l.read()
	if err != nil {
		panic(err)
	}
	l.pos++

	rr, err := l.read()
	if err != nil {
		if err == io.EOF {
			return string(r), false
		}
	}
	l.pos++

	if rr == r {
		return string(r) + string(r), true
	}

	l.unread()
	return string(r), false
}

func (l *Lexer) lexExprStart() (string, bool) {
	_, err := l.read()
	if err != nil {
		panic(err)
	}
	l.pos++

	r, err := l.read()
	if err != nil {
		if err == io.EOF {
			return string(r), false
		}
	}
	l.pos++

	if r == '(' {
		return "$(", true
	}

	l.unread()
	return "$", false
}

func (l *Lexer) lexIdent() (string, bool) {
	var token string
	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return token, true
			}
		}

		l.pos++

		if unicode.IsSpace(r) || l.oneOf(r, []rune{',', '=', '&', '|', '$', '(', ')'}) {
			l.unread()
			return token, true
		}

		token = token + string(r)
	}
}

func (l *Lexer) oneOf(rune rune, runes []rune) bool {
	for _, r := range runes {
		if r == rune {
			return true
		}
	}

	return false
}

// NewLexer Creates a new lexer instance
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    0,
		reader: bufio.NewReader(reader),
	}
}
