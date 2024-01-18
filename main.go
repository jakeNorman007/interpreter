package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
)

//this is the interpreter from the terminal, has a logo and in the main function it will error if something goes
//wrong then it prints the logo.
//finally, repl.Start does just that and takes user input and spits out at this moment the tokens and their literals
//in a struct/ key value form.
const logo = `
 ____  _  _  _       _    _
|  __|| || |(_)     | |  | |
| |_  | || | _  __  | |_ | |_
|  _| | || || |/   \| __|| __|
| |__ | || || || | || |_ | |_   _
|____||_||_||_|\___/\___|\___| |_|
`
func main() {
     _, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", logo)
    repl.Start(os.Stdin, os.Stdout)
}
