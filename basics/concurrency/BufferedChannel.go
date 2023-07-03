package main
import "fmt"


func main() {
    channel := make(chan string, 2)

    channel <- "First message"
    channel <- "Second message"

    fmt.Println(<-channel)
    fmt.Println(<-channel)
}
