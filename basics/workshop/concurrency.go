package main

import (
    "fmt"
    "time"
)

func sendData(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send data to the channel
        time.Sleep(time.Second)
    }
    close(ch) // Close the channel after sending data
}

func main() {
    ch := make(chan int)

    // Start a goroutine to send data to the channel
    go sendData(ch)

    // Receive data from the channel and print it
    for num := range ch {
        fmt.Println("Received:", num)
    }
}

