package lexer

import "fmt"

// TokenType Represents the type of a token
type TokenType int

// TokenTypes
const (
	EOF     TokenType = iota // EOF end-of-file token
	UNKNOWN                  // UNKNOWN Unrecognised token
	AND                      // AND and token
	OR                       // OR or token
	ASSIGN                   // ASSIGN assignment token
	COMMA                    // COMMA comma token
	EXPRS                    // EXPRS expression begin token
	EXPRE                    // EXPRE expression end token
	IDENT                    // IDENT identifier token
)

var tokenTypes = []string{
	EOF:     "EOF",
	UNKNOWN: "UNKNOWN",
	AND:     "AND",
	OR:      "OR",
	ASSIGN:  "ASSIGN",
	COMMA:   "COMMA",
	EXPRS:   "EXPR_S",
	EXPRE:   "EXPR_E",
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

func (t *Token) String() string {
	return "position: " + fmt.Sprint(t.Position) + "\t type:" + t.Type.String() + "\t\t value: " + t.Value
}

// IsEOF Checks if the token is an EOF token
func (t *Token) IsEOF() bool {
	return t.Type == EOF
}

// IsUnknown Checks if the current token is an UNKNOWN token
func (t *Token) IsUnknown() bool {
	return t.Type == UNKNOWN
}

// IsAnd Checks if the current token is an AND token
func (t *Token) IsAnd() bool {
	return t.Type == AND
}

// IsOr Checks if the current token is an OR token
func (t *Token) IsOr() bool {
	return t.Type == OR
}

// IsAssign Checks if the current token is an ASSIGN token
func (t *Token) IsAssign() bool {
	return t.Type == ASSIGN
}

// IsComma Checks if the current token is a COMMA token
func (t *Token) IsComma() bool {
	return t.Type == COMMA
}

// IsExprStart Checks if the current token is an EXPR_START token
func (t *Token) IsExprStart() bool {
	return t.Type == EXPRS
}

// IsExprEnd Checks if the current token is an EXPR_END token
func (t *Token) IsExprEnd() bool {
	return t.Type == EXPRE
}

// IsIdent Checks if the current token is an IIDENT token
func (t *Token) IsIdent() bool {
	return t.Type == IDENT
}
