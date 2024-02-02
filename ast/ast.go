package ast

import (
    "bytes"
    "github.com/JakeNorman007/interpreter/token"
)

//***********************AST Summary************************************
//*                                                                    *
//*  *ast.Program --> *ast.LetStatement --> Name --> *ast.Identifier   *
//*                                     \__> Value --> *ast.Expression *
//**********************************************************************

type Node interface {
    TokenLiteral() string
    String()       string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Program struct {
    Statements  []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

//String function and the implimenting of it with Expression statements, Let statements and Return statements
func (p *Program) String() string {
    var out bytes.Buffer

    for _, s := range p.Statements {
        out.WriteString(s.String())
    }

    return out.String()
}

func (ls *LetStatement) String() string {
    var out bytes.Buffer

    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")

    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }

    out.WriteString(";")

    return out.String()
}

func (rs *ReturnStatement) String() string {
    var out bytes.Buffer

    out.WriteString(rs.TokenLiteral() + " ")

    if rs.ReturnValue != nil {
        out.WriteString(rs.ReturnValue.String())
    }

    out.WriteString(";")

    return out.String()
}

func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }

    return ""
}

//Let Statement ast structure
type LetStatement struct {
    Token   token.Token //token.LET token
    Name    *Identifier
    Value   Expression 
}

func (ls * LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//Identifier ast structure
type Identifier struct {
    Token   token.Token //token.IDENT token
    Value   string
}

func (i *Identifier) expressionNode(){}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

//Return statement ast structure
type ReturnStatement struct {
    Token       token.Token //token.RETURN token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode(){}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

//Expression statement ast structure
type ExpressionStatement struct {
    Token       token.Token //the first token present in the expression
    Expression  Expression
}

func (es *ExpressionStatement) statementNode(){}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

//Integer literal struct
type IntegerLiteral struct {
    Token   token.Token
    Value   int64
}

func (il *IntegerLiteral) expressionNode(){}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

//Prefix Expressions
type PrefixExpression struct {
    Token       token.Token //prefix token, such as ! or -
    Operator    string
    Right       Expression
}

func (pe *PrefixExpression) expressionNode(){}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string { 
    var out bytes.Buffer
    out.WriteString("(")
    out.WriteString(pe.Operator)
    out.WriteString(pe.Right.String())
    out.WriteString(")")

    return out.String()
}



