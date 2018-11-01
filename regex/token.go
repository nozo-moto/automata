package regex

const (
	ILLEGAL = "ILEEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"

	LPAREN   = "("
	RPAREN   = ")"
	PLUS     = "+"
	ASTERISK = "*"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
