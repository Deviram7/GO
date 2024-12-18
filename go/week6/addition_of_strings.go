package main

import "fmt"

func main() {
    fmt.Print("enter first string:")
    var first string
    fmt.Scanln( & first)
    fmt.Print("enter second sting:")
    var second string
    fmt.Scanln( & second)
    fmt.Print(first + second)
}