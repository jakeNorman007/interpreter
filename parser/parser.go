package parser

import (
	"fmt"

	"github.com/JakeNorman007/interpreter/ast"
	"github.com/JakeNorman007/interpreter/lexer"
	"github.com/JakeNorman007/interpreter/token"
)

//parser calls the lexer and holds two values of its own like the lexer itself
//curToken is like currentToken and peepToken is like peepToken
type Parser struct {
    l *lexer.Lexer
    curToken    token.Token
    peepToken   token.Token
    errors      []string
}

//function that is a new Parser instance, it calls the next token constantly
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l, errors: []string{},}

    //reading two tokens, so the current token and peeped token are set
    p.nextToken()
    p.nextToken()

    return p
}

//this function throws the errors I give to it. At the moment it is just peepError
func (p *Parser) Errors() []string {
    return p.errors
}

//this error message is thrown when there is an iissue with peeping the next token
func (p *Parser) peepError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peepToken.Type)
    p.errors = append(p.errors, msg)
}

//function is just the parser moving to the next token
func (p *Parser) nextToken() {
    p.curToken = p.peepToken
    p.peepToken = p.l.NextToken()
}

//instance of the parser that looks at the program given to the parser
//program var references the tree that is the current program being ran through the parser, takes in an array of statements
//in a struct. Those become the tree statements
//next it loops through the tokens, which are then turned into statements, which is stored in a variable and appends them
//to the program, all until the parser feached the EOF literal
func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()

        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }

        p.nextToken()
    }
    return program
}

//parses the satements, currently only parses Let statements. Needs to be extended to take in others. If the satement
//isn't a let statement it then just reutrns nil
func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}

    if !p.expectPeep(token.IDENT) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeep(token.ASSIGN) {
        return nil
    }

    //skipping this expression until we see a semicolon
    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}

func (p *Parser) peepTokenIs(t token.TokenType) bool {
    return p.peepToken.Type == t
}

func (p *Parser) expectPeep(t token.TokenType) bool {
    if p.peepTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peepError(t)
        return false
    }
}









