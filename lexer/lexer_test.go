package lexer

import (
    "testing"
    "github.com/JakeNorman007/interpreter/token"
)

//Test case for the lexer
func TestNextToken(t *testing.T){
    input := `let five = 5;

              let ten = 10;

              let add = func(x,y) {
                x + y;
              };

              let result = add(five, ten);
              !-/*5;
              5 < 10 > 5;

              if (5 < 10) {
                return true; 
              } else {
                return false;
              }

              15 == 15;
              13 != 9;

              ` 

    tests := []struct{
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.LET, "let"},
        {token.IDENT, "five"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "ten"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "func"},
        {token.LEFTPAREN, "("},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.RIGHTPAREN, ")"},
        {token.LEFTBRACE, "{"},
        {token.IDENT, "x"},
        {token.PLUS, "+"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},
        {token.RIGHTBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "result"},
        {token.ASSIGN, "="},
        {token.IDENT, "add"},
        {token.LEFTPAREN, "("},
        {token.IDENT, "five"},
        {token.COMMA, ","},
        {token.IDENT, "ten"},
        {token.RIGHTPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.BANG, "!"},
        {token.MINUS, "-"},
        {token.SLASH, "/"},
        {token.ASTERISK, "*"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "5"},
        {token.LESSTHAN, "<"},
        {token.INT, "10"},
        {token.GREATERTHAN, ">"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.IF, "if"},
        {token.LEFTPAREN, "("},
        {token.INT, "5"},
        {token.LESSTHAN, "<"},
        {token.INT, "10"},
        {token.RIGHTPAREN, ")"},
        {token.LEFTBRACE, "{"},
        {token.RETURN, "return"},
        {token.TRUE, "true"},
        {token.SEMICOLON, ";"},
        {token.RIGHTBRACE, "}"},
        {token.ELSE, "else"},
        {token.LEFTBRACE, "{"},
        {token.RETURN, "return"},
        {token.FALSE, "false"},
        {token.SEMICOLON, ";"},
        {token.RIGHTBRACE, "}"},
        {token.INT, "15"},
        {token.EQUAL, "=="},
        {token.INT, "15"},
        {token.SEMICOLON, ";"},
        {token.INT, "13"},
        {token.NOT_EQUAL, "!="},
        {token.INT, "9"},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},

    }

    l := New(input)

    for i, tt := range tests{
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
            i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
            i, tt.expectedLiteral, tok.Literal)
        }
    }
}
