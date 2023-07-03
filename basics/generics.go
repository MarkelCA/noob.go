package main

import (
    "fmt"
    "golang.org/x/exp/constraints"
)

type MyInt int

type UserId interface {
    uint | uint8 | uint64 
}

func Add(x, y int) int {
    return x + y
}

func GenericIntAdd[T ~int](x, y T) T {
    return x + y
}

func GenericOrderedAdd[T constraints.Ordered](x, y T) T {
    return x + y
}

func GenericNumberAdd[T constraints.Integer | constraints.Float](x, y T) T {
    return x + y
}

func GenericUserIdAdd[T UserId](x, y T) T {
    return x + y
}

func main() {
    x := MyInt(1)
    y := MyInt(2)

    //r := Add(x, y)
    r := GenericIntAdd(x, y)

    fmt.Printf("Result: %v\n", r)
    fmt.Printf("Result: %v\n", GenericOrderedAdd(1.1, 6.6))
    fmt.Printf("Result: %v\n", GenericOrderedAdd("aoeu", "aoeu"))
    fmt.Printf("Result: %v\n", GenericNumberAdd(6, 5))
    //fmt.Printf("Result: %v", GenericAdd(1.1, 6))

    fmt.Printf("Result: %v\n", GenericUserIdAdd(uint(676), uint(6664)))

}
