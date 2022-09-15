package main

import (
    "fmt"
)

func PassByValue(a *[5]int) {
    a[1] = 123
} 

func PassByReference(slice []int) {
    slice[1] = 456
}

func main() {
    var a = [5]int{1, 2, 3, 4, 5}

    PassByValue(&a)
    fmt.Println("Pass By Value")
    fmt.Println("a: ", a)

    fmt.Println("Pass By Reference")
    PassByReference(a[:])
    fmt.Println("a: ", a)
}
