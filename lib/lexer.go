package lib

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

/* See parser-v3.go for the grammar */

// TokenType Represents the type of a token
type TokenType int

// TokenTypes
const (
	EOF     = iota // EOF end-of-file token
	UNKNOWN        // UNKNOWN Unrecognised token
	AND            // AND and token
	OR             // OR or token
	ASSIGN         // ASSIGN assignment token
	COMMA          // COMMA comma token
	IDENT          // IDENT identifier token
)

var tokenTypes = []string{
	EOF:     "EOF",
	UNKNOWN: "UNKNOWN",
	AND:     "AND",
	OR:      "OR",
	ASSIGN:  "ASSIGN",
	COMMA:   "COMMA",
	IDENT:   "IDENT",
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

// Token Represents a token
// Contains information about the position, type and value of the token
// Input is a single line of text, so position refers to the column position within the line
type Token struct {
	Type     TokenType
	Position int
	Value    string
}

func (t Token) String() string {
	return "{ position: " + fmt.Sprint(t.Position) + "\t type:" + t.Type.String() + "\t value: " + t.Value + " }"
}

// Lexer Represents a lexer
type Lexer struct {
	pos    int
	reader *bufio.Reader
}

// Lex Scans the input for the next token.
// It returns the position and type of the token and the literal value
func (l *Lexer) Lex() Token {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{Type: EOF, Position: l.pos, Value: ""}
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

func (l *Lexer) unread() {
	err := l.reader.UnreadRune()
	if err != nil {
		panic(err)
	}

	l.pos--
}

func (l *Lexer) lexAnd() (string, bool) {
	return l.lexConsecutive('&')
}

func (l *Lexer) lexOr() (string, bool) {
	return l.lexConsecutive('|')
}

func (l *Lexer) lexConsecutive(r rune) (string, bool) {
	_, _, err := l.reader.ReadRune()
	if err != nil {
		panic(err)
	}
	l.pos++

	rr, _, err := l.reader.ReadRune()
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

func (l *Lexer) lexIdent() (string, bool) {
	var token string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return token, true
			}
		}

		l.pos++

		if !unicode.IsSpace(r) {
			token = token + string(r)
		} else {
			l.unread()
			return token, true
		}
	}
}

// NewLexer Creates a new lexer instance
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    0,
		reader: bufio.NewReader(reader),
	}
}
