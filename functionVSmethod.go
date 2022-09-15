package main

import(
    "fmt"
)

type mytype []int

func (p *mytype) testMethod() {
    (*p)[2] = 44
}

func testFunction(p *mytype) {
    (*p)[2] = 55
}

func main() {
    p := mytype([]int{1, 2, 3})

    // The method permits to run with the reciever as a pointer
    p.testMethod()
    fmt.Printf("%v\n", p)

    // The function needs the argument passed as a reference
    testFunction(&p)
    fmt.Printf("%v\n", p)
}

