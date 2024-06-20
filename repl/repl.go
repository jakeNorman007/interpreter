package repl

import (
	"io"
	"fmt"
	"bufio"
	"github.com/JakeNorman007/interpreter/evaluator"
	"github.com/JakeNorman007/interpreter/lexer"
	"github.com/JakeNorman007/interpreter/object"
	"github.com/JakeNorman007/interpreter/parser"
    "github.com/charmbracelet/lipgloss"
)

const PROMPT = "::: "

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6"))

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()

    for {
        fmt.Printf(logoStyle.Render(PROMPT))
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParseErrors(out, p.Errors())
            continue
        }
        
        evaluated := evaluator.Eval(program, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }

}

func printParseErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t" + msg + "\n")
    }
}
