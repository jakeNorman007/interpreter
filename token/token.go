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
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LESSTHAN = "<"
    GREATERTHAN = ">"

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
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"

    /// Doubles
    EQUAL = "=="
    NOT_EQUAL = "!="
)

//bank for actual key phrases we will see in the language
var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

//function looks in the keywords table *above* to check if the current indentifier is actually a keyword.
//if yes it returns the TokenType constant of the keyword, if no it goes back to token.IDENT, the token
//type for all indetifiers we have defined
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENT
}
