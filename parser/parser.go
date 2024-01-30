package parser

import (
	"fmt"
    "strconv"
	"github.com/JakeNorman007/interpreter/ast"
	"github.com/JakeNorman007/interpreter/lexer"
	"github.com/JakeNorman007/interpreter/token"
)

const (
    _ int = iota
    LOWEST
    EQUALS          //==
    LESSGREATER     //> or <
    SUM             //+
    PRODUCT         // *
    PREFIX          //-X or !X
    CALL            //myFunction(x)
)

type Parser struct {
    l               *lexer.Lexer
    curToken        token.Token
    peepToken       token.Token
    errors          []string
    prefixParseFns  map[token.TokenType]prefixParseFn
    infixParseFns   map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l, errors: []string{},}

    p.nextToken()
    p.nextToken()

    p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
    p.registerPrefix(token.IDENT, p.parseIdentifier)
    p.registerPrefix(token.INT, p.parseIntegerLiteral)

    return p
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) peepError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peepToken.Type)
    p.errors = append(p.errors, msg)
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
    case token.RETURN:
        return p.parseReturnStatement()
    default:
        return p.parseExpressionStatement()
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

    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}

    p.nextToken()

    //TODO: skipping expressions until a semicolon is hit
    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    stmt := &ast.ExpressionStatement{Token: p.curToken}

    stmt.Expression = p.parseExpression(LOWEST)

    if p.peepTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.prefixParseFns[p.curToken.Type]
    if prefix == nil {
        return nil
    }

    leftExp := prefix()

    return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
    lit := &ast.IntegerLiteral{Token: p.curToken}

    value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
    if err != nil {
        msg := fmt.Sprintf("could not parse %q as an integer", p.curToken.Literal)
        p.errors = append(p.errors, msg)
        return nil
    }

    lit.Value = value
    
    return lit
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

//prefix, infix parsing
//maps are added in the Parser struct up top.
//prefix and infix helper methods as well, they add entries to the maps
type (
    prefixParseFn func() ast.Expression
    infixParseFn func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
    p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
    p.infixParseFns[tokenType] = fn
}




