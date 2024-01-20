package ast

import (
    "testing"
    "github.com/JakeNorman007/interpreter/token"
)

func TestString(t *testing.T) {
    program := &Program {
        Statements: []Statement {
            &LetStatement {
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier {
                    Token: token.Token{Type: token.IDENT, Literal: "variable"},
                    Value: "variable",
                },
                Value: &Identifier {
                    Token: token.Token{Type: token.IDENT, Literal: "anotherVariable"},
                    Value: "anotherVariable",
                },
            },
        },
    }

    if program.String() != "let variable = anotherVariable;" {
        t.Errorf("program.String() is wrong, got=%q", program.String())
    }
}
