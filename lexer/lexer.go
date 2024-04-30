package lexer

import "github.com/JakeNorman007/interpreter/token"

type Lexer struct {
    input           string //as name suggests this is the input itself 
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

    l.eatWhitespace()

    switch l.ch {
    case '=':
        if l.peepChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.ch)}
        } else {
            tok = newToken(token.ASSIGN, l.ch)
        }

    case '!':
        if l.peepChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = token.Token{Type: token.NOT_EQUAL, Literal: string(ch) + string(l.ch)}
        } else {
            tok = newToken(token.BANG, l.ch)
        }

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
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '*':
        tok = newToken(token.ASTERISK, l.ch)
    case '/':
        tok = newToken(token.SLASH, l.ch)
    case '<':
        tok = newToken(token.LESSTHAN, l.ch)
    case '>':
        tok = newToken(token.GREATERTHAN, l.ch)
    case '"':
        tok.Type = token.STRING
        tok.Literal = l.readString()
    case '[':
        tok = newToken(token.LEFTBRACKET, l.ch)
    case ']':
        tok = newToken(token.RIGHTBRACKET, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
           tok.Type = token.INT
           tok.Literal = l.readNumber()
           return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readString() string {
    position := l.position + 1
    for {
        l.readChar()

        if l.ch == '"' || l.ch == 0 {
            break
        }
    }

    return l.input[position:l.position]
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
    //TODO: add other number types. Floats, hexadecimal, octs. For now it's not supported. Will add later!!
}

func (l *Lexer) eatWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
    position := l.position

    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) peepChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}
