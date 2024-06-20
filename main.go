package main

import (
    "os"
    "fmt"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
    "github.com/charmbracelet/lipgloss"
)

const logo = `
**********************************************
* EEEEEE EE    EE    EE EEEEEE EEEEEE EEEEEE *
* EE     EE    EE    EE EE  EE   EE     EE   *
* EEEE   EE    EE    EE EE  EE   EE     EE   *
* EE     EE    EE    EE EE  EE   EE     EE   *
* EEEEEE EEEEE EEEEE EE EEEEEE   EE     EE   * 
**********************************************
An interpreter for the Elliott programming language.

Use Ctrl+C to exit.
`
var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6"))

func main() {
     _, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", logoStyle.Render(logo))
    repl.Start(os.Stdin, os.Stdout)
}
