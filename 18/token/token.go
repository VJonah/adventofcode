package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	EOF      = "EOF"
	LPAREN   = "("
	RPAREN   = ")"
	INT      = "INT"
	ASTERISK = "*"
	PLUS     = "+"
)
