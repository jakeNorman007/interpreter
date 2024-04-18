package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
)

const logo = `
 ____  _
|  __|| |
| |_  | |
|  _| | |
| |__ | | _
|____||_||_|
`
func main() {
     _, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", logo)
    repl.Start(os.Stdin, os.Stdout)
}
