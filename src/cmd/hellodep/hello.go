package main

import "fmt"

import "internal/pkg1"
//import "internal/commons"

func main() {
    fmt.Println("hello world ")
    text := pkg1.GetText()
    fmt.Println("pkg1.GetText() : " + text)
}
