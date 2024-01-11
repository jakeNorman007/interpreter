package parser

import (
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
}

//function that is a new Parser instance, it calls the next token constantly
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    //reading two tokens, so the current token and peeped token are set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peepToken
    p.peepToken = p.l.NextToken()
}

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
        return false
    }
}









