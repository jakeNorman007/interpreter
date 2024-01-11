package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
)

//not the actual logo
const logo = `
 ____  _  _  _       _    _
|  __|| || |(_)     | |  | |
| |_  | || | _  __  | |_ | |_
|  _| | || || |/   \| __|| __|
| |__ | || || || | || |_ | |_   _
|____||_||_||_|\___/\___|\___| |_|
`
func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Welcome %s, to the Elliott programming language.", user.Username)
    fmt.Printf("%s\n", logo)
    repl.Start(os.Stdin, os.Stdout)
}
