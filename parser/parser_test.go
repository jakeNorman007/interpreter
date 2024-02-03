package parser

import (
	"fmt"
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

    if len(program.Statements) != 1 {
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

    if ident.Value != "foobarbaz" {
        t.Errorf("ident.Value is not %s, got=%s", "foobarbazz", ident.Value)
    }

    if ident.TokenLiteral() != "foobarbaz" {
        t.Errorf("ident.TokenLiteral is not %s, got=%s", "foobarbazz", ident.TokenLiteral())
    }
}

//parser test for integer literals ex. 5; 10; 55; etc.
func TestIntegerLiteral(t *testing.T) {
    input := "6;"

    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not as ast.ExpressionStatement, got=%T", program.Statements[0])
    }

    literal, ok := stmt.Expression.(*ast.IntegerLiteral)
    if !ok {
        t.Fatalf("exp not *ast.IntegerLiteral, got=%T", stmt.Expression)
    }
    
    if literal.Value != 6 {
        t.Errorf("literal.Value not %d, got=%d", 6, literal.Value)
    }

    if literal.TokenLiteral() != "6" {
        t.Errorf("literal.TokenLiteral is not %s, got=%s", "6", literal.TokenLiteral())
    }
}

//parser test for prefix expressions - or !
func TestParsingPrefixExpressions(t *testing.T) {
    prefixTests := []struct {
        input           string
        operator        string
        integerValue    int64
    }{
        {"!5;", "!", 5},
        {"-15", "-", 15},
    }

    for _, tt := range prefixTests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 1 {
            t.Fatalf("program does not contain %d statements, got=%d", 1, len(program.Statements))
        }

        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statement[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.PrefixExpression)
        if !ok {
            t.Fatalf("stmt in not ast.PrefixExpression, got=%T", stmt.Expression)
        }
        
        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator is not %s, got=%s", tt.operator, exp.Operator)
        }

        if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
            return
        }
    }
}

//testIntegerLiteral helper function to use with PrefixExpression test
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
    integ, ok := il.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("il not *ast.IntegerLiteral, got=%T", il)
        return false
    }

    if integ.Value != value {
        t.Errorf("integ.Value is not %d, got=%d", value, integ.Value)
        return false
    }

    if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
        t.Errorf("integ.TokenLiteral is not %d, got=%s", value, integ.TokenLiteral())
        return false
    }

    return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
    ident, ok := exp.(*ast.Identifier)
    if !ok {
        t.Errorf("exp is not *ast.Identifier, got=%T", exp)
        return false
    }

    if ident.Value != value {
        t.Errorf("ident.Value is not %s, got=%s", value, ident.Value)
        return false
    }

    if ident.TokenLiteral() != value {
        t.Errorf("ident.TokenLiteral is not %s, got=%s", value, ident.TokenLiteral())
        return false
    }

    return true
}

//Infix expressions test
func TestParsingInfixExpressions(t *testing.T) {
    infixTests := []struct {
        input           string
        leftValue       int64
        operator        string
        rightValue      int64
    }{
        {"5 + 5;", 5, "+", 5},
        {"5 - 5;", 5, "-", 5},
        {"5 * 5;", 5, "*", 5},
        {"5 / 5;", 5, "/", 5},
        {"5 > 5;", 5, ">", 5},
        {"5 < 5;", 5, "<", 5},
        {"5 == 5;", 5, "==", 5},
        {"5 != 5;", 5, "!=", 5},
    }

    for _, tt := range infixTests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 1 {
            t.Fatalf("program.Statements does not contain %d statements, got=%d\n", 1, len(program.Statements))
        }

        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.InfixExpression)
        if !ok {
            t.Fatalf("stmt.Expression is not ast.Expression, got=%T", stmt.Expression)
        }

        if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
            return
        }

        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator in not %s, got=%s", tt.operator, exp.Operator)
        }

        if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
            return
        }
    }
}

//Operator precedence test
func TestOperatorPrecedenceParsing(t *testing.T) {
    tests := []struct {
        input       string
        expected    string
    }{
        {
            "-a * b",
            "((-a) * b)",
        },
        {
            "!-a",
            "(!(-a))",
        },
        {
            "a + b - c",
            "((a + b) - c)",
        },
        {
            "a - b - c",
            "((a - b) - c)",
        },
        {
            "a * b * c",
            "((a * b) * c)",
        },
        {
            "a * b / c",
            "((a * b) / c)",
        },
        {
            "a + b / c",
            "(a + (b / c))",
        },
        {
            "a + b * c + d / e - f",
            "(((a + (b * c)) + (d / e)) - f)",
        },
        {
            "3 + 4; -5 * 5",
            "(3 + 4)((-5) * 5)",
        },
        {
            "5 > 4 != 3 > 4",
            "((5 > 4) != (3 > 4))",
        },
        {
            "5 < 4 != 3 > 4",
            "((5 < 4) != (3 > 4))",
        },
        {
            "3 + 4 * 5 == 3 * 1 + 4 * 5",
            "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
        },
        {
            "3 + 4 * 5 == 3 * 1 + 4 * 5",
            "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
        },
    }
    
    for _, tt := range tests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        actual := program.String()

        if actual != tt.expected {
            t.Errorf("expected=%q, got=%q", tt.expected, actual)
        }
    }
}






