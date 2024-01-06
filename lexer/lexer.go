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

    //in case we come across whitespace, which shouldn't throw an error this function skips it when
    //it's being read by the lexer.
    l.eatWhitespace()

    switch l.ch {
    case '=':
        if l.peepChar() == '=' {
            //saving l.ch in a local variable so when we check the next character I don't lose the first
            //i.e. saves = so when it peeps =, the first = isn't overwritten by =, and becomes the assignment
            //token
            ch := l.ch
            l.readChar()
            tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.ch)}
        } else {
            tok = newToken(token.ASSIGN, l.ch)
        }

    case '!':
        if l.peepChar() == '=' {
            //saving l.ch in a local variable so when we check the next character I don't lose the first
            //i.e. saves ! so when it peeps = ! isn't overwritten by =
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

//similar to readIndetifier, but it will read if it's a number
func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

//similar to what is in the above comment, check if what the lexer is reading is actually a number
//basically the number version of isLetter
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
    //TODO: add other number types. Floats, hexadecimal, octs. For now it's not supported. Will add later!!
}

//eat whitespace function that ignores a read if it's a whitespace. In doing this it will not throw an error
//and the lexer will just move past it to the next character.
//includes not only whitespace, but regex returns as well
func (l *Lexer) eatWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
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

//peepChar i s a function that works like readChar, but does not incriment the position being read, instead it just looks ahead
//this helps with things like == for equal to or != for not equal to.
//if it peeps ahead and say it is ==, then it will know based off the token that it is lexing == and will then read it as so
func (l *Lexer) peepChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}
