package main

import (
    "fmt"
)

type Walker interface {
    walk()
}

func Walk(w Walker) {
    fmt.Println("Walking...")
}


type Animal struct {}

type Supertype struct {
    w Walker
}

func (a Animal) walk() {
    fmt.Println("Walking...")
}

func main() {
    w := Animal{}

    s:= Supertype{w}

    fmt.Println(s)

    Walk(w)
}

// The following function is illegal (https://stackoverflow.com/questions/65440847/invalid-receiver-for-pointer-alias-type):
//func (w Walker) walkV2() {
    //fmt.Println("Walking v2...")
//}

// Pointers & Interfaces as receiver types - Why the design choice. (ChatGPT answer)
//The design choice of disallowing pointers and interfaces as receiver base types in Go was made for several reasons.

//Firstly, it helps to simplify the language and make it more consistent. By disallowing pointer and interface types as receiver base types, Go avoids the need to deal with issues related to indirection and type conversion, which can complicate the language and make it harder to use.

//Secondly, disallowing pointers and interfaces as receiver base types encourages developers to use value types instead. This can help to improve performance, as value types are typically more efficient than pointer types, especially for small objects.

//Finally, requiring the receiver type and the method to be defined in the same package helps to promote encapsulation and maintainability. By restricting methods to operate only on objects defined within the same package, Go encourages developers to keep related code together, which can make it easier to understand and modify.

//Overall, these design choices reflect Go's emphasis on simplicity, performance, and maintainability, and they have helped to make the language a popular choice for a wide range of applications.
