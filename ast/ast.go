package ast

import "github.com/JakeNorman007/interpreter/token"

//interfaces for our ast, every node in the ast has to call the Node interface, hence why
//in Statement and Expression it's called. The Node then has to return a token literal,
//so in Statement and in Expression they will get called
type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

//Program node serves as the root node for every tree our parser generates
type Program struct {
    Statements  []Statement
}

//TokenLiteral points to the current program being ran p. If the length of the statements in the
//program is greater than 0, then it returns the Staement as the correct Token Literal. If not
//it returns nothing.
func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

type LetStatement struct {
    Token   token.Token //token.LET token
    Name    *Identifier
    Value   Expression 
}

func (ls * LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
    Token   token.Token //token.IDENT token
    Value   string
}

func (i *Identifier) ExpressionNode(){}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//**AST Summary**
//
//  *ast.Program --> *ast.LetStatement --> Name --> *ast.Identifier
//                                     \__> Value --> *ast.Expression


