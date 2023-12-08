package token

type TokenType string

type Token struct{
    Type    TokenType
    Literal string
}

const(
    // Sorts of errors
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers, literals
    IDENT = "IDENT" // add, foobar, x, y
    INT = "INT" // integers... 1, 2, 3, 4, etc..

    // Ops
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LEFTPAREN = "("
    RIGHTPAREN = ")"
    LEFTBRACE = "{"
    RIGHTBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)
