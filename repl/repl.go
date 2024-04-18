package repl

import (
    "bufio"
    "fmt"
    "io"
    "github.com/JakeNorman007/interpreter/lexer"
    "github.com/JakeNorman007/interpreter/parser"
    "github.com/JakeNorman007/interpreter/evaluator"
)

const PROMPT = "[el]>> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
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
        
        evaluated := evaluator.Eval(program)
        if evaluated != nil {
            io.WriteString(out, program.String())
            io.WriteString(out, "\n")
        }
    }

}

func printParseErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t" + msg + "\n")
    }
}
