package main

import (
    "fmt"
)

type IFace interface {
    SetSomeField(newValue string)
    GetSomeField() string
}

type Implementation struct {
    someField string
}

func (i Implementation) GetSomeField() string {
    return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
    i.someField = newValue
}

func Create() *Implementation {
    obj := &Implementation{someField: "Hello"}
    return obj // <= Offending line
}

func main() {
    a := Create()
    a.SetSomeField("World")
    fmt.Println(a.GetSomeField())
}
