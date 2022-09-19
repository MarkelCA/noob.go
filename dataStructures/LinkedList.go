package main
import "fmt"

type node struct {
    value int
    next *node
}

type linkedList struct {
    head *node
}

func (list *linkedList) prepend(n node) {
    second := list.head
    list.head = &n
    list.head.next = second
}

func (list linkedList) printListData() {
    currentNode := list.head
    for currentNode != nil {
        fmt.Printf("%d -> ", currentNode.value)
        currentNode = currentNode.next
    }
}

func (list *linkedList) deleteWithValue(value int) {
    currentNode := list.head

    if value == list.head.value {
        list.head = list.head.next
    }
    for currentNode != nil {
        if currentNode.next != nil && currentNode.next.value == value {
            currentNode.next = currentNode.next.next
        }
        currentNode = currentNode.next
    }
}

func main() {
    n1 := node{value:1}
    n2 := node{value:2}
    n3 := node{value:3}

    list := linkedList {}

    list.prepend(n1)
    list.prepend(n2)
    list.prepend(n3)
    
    list.printListData()

    list.deleteWithValue(2)

    fmt.Println()

    list.printListData()
}
