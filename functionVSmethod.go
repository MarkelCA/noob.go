package main

import("fmt")

type myArray [3]int


func (p *myArray) testMethod() {
    (*p)[2] = 44
}

func testFunction(p *myArray) {
    (*p)[2] = 55
}

func main() {
    p := myArray([3]int{1, 2, 3})

    // The method permits to run with the reciever as a pointer
    p.testMethod()
    fmt.Printf("%v\n", p)

    // The function needs the argument passed as a reference
    testFunction(&p)
    fmt.Printf("%v\n", p)
}

