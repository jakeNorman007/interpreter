package parser

import (
    "testing"
    "github.com/JakeNorman007/interpreter/ast"
    "github.com/JakeNorman007/interpreter/lexer"
)

//parser test for let statement currently
func TestLetStatement(t *testing.T) {
    input :=`
            let x = 5;
            let y = 10;
            let foo = 838383;
            `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if program == nil {
        t.Fatalf("ParseProgram() retutned nil")
    }
    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foo"},
    }

    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier){
            return
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()

    if len(errors) == 0 {
        return
    }

    t.Errorf("parser encountered %d errors.", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }

    t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral is not 'let', got=%q", s.TokenLiteral())
        return false
    }
    
    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s is not *ast.Statement, got=%T", s)
        return false
    }

    if letStmt.Name.Value != name {
        t.Errorf("letStmt.Name.Value is not %s, got=%s", name, letStmt.Name.Value)
        return false
    }

    if letStmt.Name.TokenLiteral() != name {
        t.Errorf("s.Name is nor %s, got=%s", name, letStmt.Name)
        return false
    }

    return true
}

//parser test for return statements
func TestReturnStatements(t *testing.T) {
    input := `
            return 5;
            return 14;
            return 394857584;
            `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    for _, stmt := range program.Statements {
        returnStmt, ok := stmt.(*ast.ReturnStatement)
        if !ok {
            t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
            continue
        }
        if returnStmt.TokenLiteral() != "return" {
            t.Errorf("returnStmt.TokenLiteral is not 'return', got=%q", returnStmt.TokenLiteral())
        }
    }
}

//parser test for identifier expressions
func TestIdentifierExpression(t *testing.T) {
    input := `foobarbaz;`

        l := lexer.New(input)
        p := New(l)

        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 3 {
            t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
        }

        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
        }
        
        ident, ok := stmt.Expression.(*ast.Identifier)
        if !ok {
            t.Fatalf("expression is not an *ast.Identifier, got=%T", stmt.Expression)
        }

        if ident.Value != "foobarbazz" {
            t.Errorf("ident.Value is not %s, got=%s", "foobarbazz", ident.Value)
        }

        if ident.TokenLiteral() != "foobarbazz" {
            t.Errorf("ident.TokenLiteral is not %s, got=%s", "foobarbazz", ident.TokenLiteral())
        }
}
































