// Channel example.
// In this example the channel can only be used for one iteration
package main

import (
    "fmt"
    "time"
)

func main() {
    evilNinjas := []string{"Tommy", "Jhonny", "Bobby", "Andy"}

    smokeSignal := make(chan bool)
    for _, ninja := range evilNinjas {
        go attack(ninja, smokeSignal)
        //smokeSignal <- false # Triggers deadlock Error
    }

    fmt.Printf("Attacked: %v",<-smokeSignal)

}

func attack(ninja string, attacked chan bool) {
    time.Sleep(time.Second)
    fmt.Println("Throwing ninja stars to", ninja)
    attacked <- true

}
