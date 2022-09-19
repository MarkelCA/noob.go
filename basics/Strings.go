package main

import("fmt")

func main() {
    s := "This is my string"

    s[1] = byte(4)
    fmt.Printf("%c", s[0])
}
