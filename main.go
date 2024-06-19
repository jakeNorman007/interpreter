package main

import (
    "os"
    "fmt"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
)

const logo = `
**********************************************
* EEEEEE EE    EE    EE EEEEEE EEEEEE EEEEEE *
* EE     EE    EE    EE EE  EE   EE     EE   *
* EEEE   EE    EE    EE EE  EE   EE     EE   *
* EE     EE    EE    EE EE  EE   EE     EE   *
* EEEEEE EEEEE EEEEE EE EEEEEE   EE     EE   * 
**********************************************
`

func main() {
     _, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", logo)
    repl.Start(os.Stdin, os.Stdout)
}
