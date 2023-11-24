package lexer

import (
	"testing"

	"github.com/JakeNorman007/interpreter/token"
)

func TestNextToken(t *testing.T){
    input := `=+(){},;`

    tests := []struct{
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LEFTPAREN, "("},
        {token.RIGHTPAREN, ")"},
        {token.LEFTBRACE, "{"},
        {token.RIGHTBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt := range tests{
        tok := l.NextToken()

        if tok.Type != tt.expectedType{
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
                i, tt.expectedType, tok.Type)
        }

        if tok.Type != tt.expectedLiteral{
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
                i, tt.expectedLiteral, tok.Literal)
        }
    }
}
