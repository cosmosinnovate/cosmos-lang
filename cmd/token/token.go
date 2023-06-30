package token

/**
 * The token package defines constants representing the lexical tokens of the Cosmos-Lang programming language.
 * These constants are returned by the Scanner.NextToken() method.
 * The token package also defines a Token struct that contains a token type and a literal value.
 * The token package also defines a TokenType string type that represents a token type.
 * The token package also defines constants for the different token types.
 */

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // unknown token
	EOF     = "EOF"     // end of file

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1234567890
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	LT = "<"
	GT = ">"
	EQ = "=="
	NE = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	PERIOD    = "."
	ELLIPSIS  = "..."
	QUESTION  = "?"
	AMPERSAND = "&"
	PIPE      = "|"
	CARET     = "^"
	PLUS_EQ   = "+="
	MINUS_EQ  = "-="
	ASTER_EQ  = "*="

	// Brackets
	LPAREN     = "("
	RPAREN     = ")"
	LBRACE     = "{"
	RBRACE     = "}"
	LBRACKET   = "["
	RBRACKET   = "]"
	LT_EQ      = "<="
	GT_EQ      = ">="
	DOUBLE_EQ  = "=="
	NOT_EQ     = "!="
	DOUBLE_AND = "&&"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
