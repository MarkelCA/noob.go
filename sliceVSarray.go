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

    var a = [5]int{3, 11, 6, 20, 9}

    PassByValue(&a)
    fmt.Println("Pass By Value")
    fmt.Println("a: ", a)

    fmt.Println("Pass By Reference")
    PassByReference(a[:])
    fmt.Println("a: ", a)

}
  
