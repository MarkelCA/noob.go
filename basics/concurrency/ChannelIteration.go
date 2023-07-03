package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    channel := make(chan string)
    numRounds := 3
    go throwingNinjastar(channel, numRounds)

    for i := 0 ; i < numRounds ; i++ {
        fmt.Println(<-channel)
    }
}


func throwingNinjastar(channel chan string, numRounds int) {
    rand.Seed(time.Now().UnixNano())
    for i := 0 ; i < numRounds ; i++ {
        score := rand.Intn(10)
        channel <- fmt.Sprint("You scored:", score)
    }
}
