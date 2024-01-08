package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/JakeNorman007/interpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Welcome %s\n", user.Username)
    fmt.Printf("Elliott programming languag:\n")
    repl.Start(os.Stdin, os.Stdout)
}
