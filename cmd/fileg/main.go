package main

import (
    "fmt"
    "github.com/nkxxll/filegloss"
)

func main() {
    f := filegloss.New("../../fixtures/helloworld")
    fmt.Print(f.InfoString())
}
