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
    PRODUCT         //*
    PREFIX          //-X or !X
    CALL            //myFunctionName(x)
    INDEX
)

type Parser struct {
    l               *lexer.Lexer
    curToken        token.Token
    peepToken       token.Token
    errors          []string
    prefixParseFns  map[token.TokenType]prefixParseFn
    infixParseFns   map[token.TokenType]infixParseFn
}

var precedences = map[token.TokenType]int {
    token.EQUAL:        EQUALS,
    token.NOT_EQUAL:    EQUALS,
    token.LESSTHAN:     LESSGREATER,
    token.GREATERTHAN:  LESSGREATER,
    token.PLUS:         SUM,
    token.MINUS:        SUM,
    token.SLASH:        PRODUCT,
    token.ASTERISK:     PRODUCT,
    token.LEFTPAREN:    CALL,
    token.LEFTBRACKET:  INDEX,
}

func (p *Parser) peepPrecedence() int {
    if p, ok := precedences[p.peepToken.Type]; ok {
        return p
    }

    return LOWEST
}

func (p *Parser) curPrecedence() int {
    if p, ok := precedences[p.curToken.Type]; ok {
        return p
    }

    return LOWEST
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l, errors: []string{},}

    p.nextToken()
    p.nextToken()

    p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
    p.registerPrefix(token.IDENT, p.parseIdentifier)
    p.registerPrefix(token.INT, p.parseIntegerLiteral)
    p.registerPrefix(token.BANG, p.parsePrefixExpression)
    p.registerPrefix(token.MINUS, p.parsePrefixExpression)
    p.registerPrefix(token.TRUE, p.parseBoolean)
    p.registerPrefix(token.FALSE, p.parseBoolean)
    p.registerPrefix(token.LEFTPAREN, p.parseGroupedExpression)
    p.registerPrefix(token.IF, p.parseIfExpression)
    p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
    p.registerPrefix(token.STRING, p.parseStringLiteral)
    p.registerPrefix(token.LEFTBRACKET, p.parseArrayLiteral)
    p.registerPrefix(token.LEFTBRACE, p.parseHashLiteral)

    p.infixParseFns = make(map[token.TokenType]infixParseFn)
    p.registerInfix(token.PLUS, p.parseInfixExpression)
    p.registerInfix(token.MINUS, p.parseInfixExpression)
    p.registerInfix(token.SLASH, p.parseInfixExpression)
    p.registerInfix(token.ASTERISK, p.parseInfixExpression)
    p.registerInfix(token.EQUAL, p.parseInfixExpression)
    p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)
    p.registerInfix(token.LESSTHAN, p.parseInfixExpression)
    p.registerInfix(token.GREATERTHAN, p.parseInfixExpression)
    p.registerInfix(token.LEFTPAREN, p.parseCallExpression)
    p.registerInfix(token.LEFTBRACKET, p.parseIndexExpression)

    return p
}

func (p *Parser) parseHashLiteral() ast.Expression {
    hash := &ast.HashLiteral{Token: p.curToken}
    hash.Pairs = make(map[ast.Expression]ast.Expression)

    for !p.peepTokenIs(token.RIGHTBRACE) {
        p.nextToken()
        key := p.parseExpression(LOWEST)

        if !p.expectPeep(token.COLON) {
            return nil
        }

        p.nextToken()
        value := p.parseExpression(LOWEST)

        hash.Pairs[key] = value

        if !p.peepTokenIs(token.RIGHTBRACE) && !p.expectPeep(token.COMMA) {
            return nil
        }

    }

    if !p.expectPeep(token.RIGHTBRACE) {
        return nil
    }

    return hash
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
    exp := &ast.IndexExpression{Token: p.curToken, Left: left}
    
    p.nextToken()
    exp.Index = p.parseExpression(LOWEST)

    if !p.expectPeep(token.RIGHTBRACKET) {
        return nil
    }

    return exp
}

func (p *Parser) parseArrayLiteral() ast.Expression {
    array := &ast.ArrayLiteral{Token: p.curToken}

    array.Elements = p.parseExpressionList(token.RIGHTBRACKET)

    return array
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
    list := []ast.Expression{}

    if p.peepTokenIs(end) {
        p.nextToken()
        return list
    }

    p.nextToken()
    list = append(list, p.parseExpression(LOWEST))

    for p.peepTokenIs(token.COMMA) {
        p.nextToken()
        p.nextToken()
        list = append(list, p.parseExpression(LOWEST))
    }

    if !p.expectPeep(end) {
        return nil
    }

    return list
}

func (p *Parser) parseStringLiteral() ast.Expression {
    return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
    exp := &ast.CallExpression{Token: p.curToken, Function: function}
    exp.Arguments = p.parseExpressionList(token.RIGHTPAREN)
    return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
    args := []ast.Expression{}

    if p.peepTokenIs(token.RIGHTPAREN) {
        p.nextToken()
        return args
    }

    p.nextToken()
    args = append(args, p.parseExpression(LOWEST))

    for p.peepTokenIs(token.COMMA) {
        p.nextToken()
        p.nextToken()
        args = append(args, p.parseExpression(LOWEST))
    }

    if !p.expectPeep(token.RIGHTPAREN) {
        return nil
    }

    return args
}

func (p *Parser) parseGroupedExpression() ast.Expression {
    p.nextToken()

    exp := p.parseExpression(LOWEST)

    if !p.expectPeep(token.RIGHTPAREN) {
        return nil
    }

    return exp
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    expression := &ast.InfixExpression {
        Token:      p.curToken,
        Operator:   p.curToken.Literal,
        Left:       left,
    }

    precedence := p.curPrecedence()
    p.nextToken()
    expression.Right = p.parseExpression(precedence)

    return expression
}

func (p *Parser) parseBoolean() ast.Expression {
    return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
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

    p.nextToken()

    stmt.Value = p.parseExpression(LOWEST)

    if p.peepTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}

    p.nextToken()

    stmt.ReturnValue = p.parseExpression(LOWEST)

    if p.peepTokenIs(token.SEMICOLON) {
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
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
    }

    leftExp := prefix()

    for !p.peepTokenIs(token.SEMICOLON) && precedence < p.peepPrecedence() {
        infix := p.infixParseFns[p.peepToken.Type]
        if infix == nil {
            return leftExp
        }

        p.nextToken()

        leftExp = infix(leftExp)
    }

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

func (p *Parser) parseIfExpression() ast.Expression {
    expression := &ast.IfExpression{Token: p.curToken}

    if !p.expectPeep(token.LEFTPAREN) {
        return nil
    }

    p.nextToken()
    expression.Condition = p.parseExpression(LOWEST)

    if !p.expectPeep(token.RIGHTPAREN) {
        return nil
    }

    if !p.expectPeep(token.LEFTBRACE) {
        return nil
    }

    expression.Consequence = p.parseBlockStatement()

    if p.peepTokenIs(token.ELSE) {
        p.nextToken()

        if !p.expectPeep(token.LEFTBRACE) {
            return nil
        }

        expression.Alternative = p.parseBlockStatement()
    }

    return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
    block := &ast.BlockStatement{Token: p.curToken}
    block.Statements = []ast.Statement{}

    p.nextToken()

    for !p.curTokenIs(token.RIGHTBRACE) && !p.curTokenIs(token.EOF) {
        stmt := p.parseStatement()
        if stmt != nil {
            block.Statements = append(block.Statements, stmt)
        }

        p.nextToken()
    }

    return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
    lit := &ast.FunctionLiteral{Token: p.curToken}

    if !p.expectPeep(token.LEFTPAREN) {
        return nil
    }

    lit.Parameters = p.parseFunctionParameters()

    if !p.expectPeep(token.LEFTBRACE) {
        return nil
    }

    lit.Body = p.parseBlockStatement()

    return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
    identifiers := []*ast.Identifier{}

    if p.peepTokenIs(token.RIGHTPAREN) {
        p.nextToken()
        return identifiers
    }

    p.nextToken()

    ident := &ast.Identifier{ Token: p.curToken, Value: p.curToken.Literal }
    identifiers = append(identifiers, ident)

    for p.peepTokenIs(token.COMMA) {
        p.nextToken()
        p.nextToken()
        ident := &ast.Identifier{ Token: p.curToken, Value: p.curToken.Literal }
        identifiers = append(identifiers, ident)
    }

    if !p.expectPeep(token.RIGHTPAREN) {
        return nil
    }

    return identifiers
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

func (p *Parser) parsePrefixExpression() ast.Expression {
    expression := &ast.PrefixExpression {
        Token: p.curToken,
        Operator:   p.curToken.Literal,
    }

    p.nextToken()
    expression.Right = p.parseExpression(PREFIX)

    return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
    msg := fmt.Sprintf("no prefix parse function for %s found", t)
    p.errors = append(p.errors, msg)
}

