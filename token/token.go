package token

const (
	ILLEGAL = "ILLEGAL"
	EOF	= "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN	= "="
	PLUS	= "+"

	// Delimiters
	COMMA 	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET 	 = "LET"
)

var keywords = map[string]TokenType {
	"fn":	FUNCTION,
	"let":	LET,
}

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

func GetTokenTypeFromLiteral(literal string) TokenType {
	if tokenType, ok := keywords[literal]; ok {
		return tokenType
	}
	return IDENT
}

