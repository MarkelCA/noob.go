package main

import (
    "fmt"
    "time"
)

func main() {
    evilNinjas := []string{"Tommy", "Jhonny", "Bobby", "Andy"}

    for _, ninja := range evilNinjas {
        go attack(ninja)
    }

    time.Sleep(2 * time.Second)
}

func attack(ninja string) {
    time.Sleep(time.Second)
    fmt.Println("Throwing ninja stars to", ninja)
}
