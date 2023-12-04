package lexer

import "github.com/JakeNorman007/interpreter/token"

type Lexer struct {
    input           string 
    position        int //current position of the input 
    readPosition    int //current reading position of the input (the character following the position) 
    ch              byte //this is the current character that is being read
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    }else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LEFTPAREN, l.ch)
    case ')':
        tok = newToken(token.RIGHTPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '{':
        tok = newToken(token.LEFTBRACE, l.ch)
    case '}':
        tok = newToken(token.RIGHTBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}


