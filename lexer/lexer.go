package lexer

import "github.com/JakeNorman007/interpreter/token"

type Lexer struct {
    input           string //as name suggests this is the input itself 
    position        int //current position of the input 
    readPosition    int //current reading position of the input (the character following the position) 
    ch              byte //this is the current character that is being read
}

//function that essentially is the lexer, takes input, reads character, returns character
func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

//read character function that first checks if it is at the end of the input, if so l.ch = 0.
//fun fact 0 is ASCII for NULL
//if not at the end of the input it sets the current character as the next character of the input,
//goes on until it reads the input as 0 or End of file in the case of this interpreter
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    }else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

//NextToken is a switch that gives all of the possible tokens to be read currently, where tok is the 
//variable the token is stored at.
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

    //this is just a check that the token being read is a token, or keyword we have in the lexer,
    //a valid token if you will, if not it will throw token.ILLEGAL. Sort of error handling
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

//newToken initializes the tokens, see NextToken cases to connect the dots, takes in a token type, followed by
//the cooresponding literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

//this function takes in the indentifier, reads it and moves forward a position in the input until it comes
//across a character that is not a letter
func (l *Lexer) readIdentifier() string {
    position := l.position

    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

//as the name says checks that the argument given is actually a letter, checking for _, means we treat it as a
//letter, so for example we can use variable names such as whoa_dude and other snake case type things.
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

