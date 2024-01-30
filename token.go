
package token

type TokenType string

type Token struct {

  Type TokenType
  Literal string

}


const (
  ILLEGAL = "ILLEGAL"

  EOF = "EOF"

  IDENT = "IDENT"

  INT = "INT"

  ASSIGN = "="

  PLUS = "+"

  // Delimitters

  COMMA = ","

  SEMICOLON = ";"

  LPAREN = "("

  RPAREN = ")"

  LBRACE = "}"

  RBRACE = "}"

  //KEYWORDS

  FUNCTION = "FUNCTION"

  LET = "LET"
)


)
