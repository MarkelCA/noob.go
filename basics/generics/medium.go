package main

import (
    "fmt"
    "golang.org/x/exp/constraints"
)

type CustomData interface {
    constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
    ID int
    Name string
    Data T
}

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
    var newValues []T
    for _, v := range  values {
        newValue := mapFunc(v)
        newValues = append(newValues, newValue)
    }
    return newValues
}

type CustomMap[T comparable, V int | string] map[T]V // comparable it's any type wich can be compared to itself with the == operator

func main() {

    result := MapValues([]float64{1.1, 2.2, 3.3}, func(n float64) float64 {
        return n * 2
    })

    fmt.Printf("result: %v\n", result)

    u := User[string]{
        ID: 0,
        Name: "Markel",
        Data: "My custom data",
    }

    fmt.Println(u)

    m := make(CustomMap[int, string])
    m[3] = "blah"

    fmt.Printf("CustomMap: %v\n", m)


}
