package main

import "fmt"

type node struct {
    value int
    next *node
}
type linkedlist struct {
    head *node
}

func (l *linkedlist) prependValue(val int) {
    second := l.head
    l.head = &node{value:val, next:second}
}

func (l *linkedlist) prependAll(vals ...int) {
    for _,value := range vals {
        l.prependValue(value)
    }
}

func (l linkedlist) printList() {
    currentNode := l.head
    for currentNode != nil {
        fmt.Printf("%d -> ", currentNode.value)
        currentNode = currentNode.next
    }
}

func (l *linkedlist) deleteWithValue(value int, deleteAll bool) {
    // Remove in the head
    if l.head.value == value {
        l.head = l.head.next
        if !deleteAll {
            return
        }
    }

    // Remove in the body
    currentNode := l.head
    for currentNode != nil {
        if currentNode.next != nil && currentNode.next.value == value {
            currentNode.next = currentNode.next.next 
            if !deleteAll {
                return
            }
        }
        currentNode = currentNode.next
    }
}

func (l linkedlist) getMiddle() (int,int) {
    slow_ptr := l.head
    fast_ptr := l.head

    i := 0
    for fast_ptr != nil && fast_ptr.next != nil {
        slow_ptr = slow_ptr.next
        fast_ptr = fast_ptr.next.next
        i++
    }

    return i, slow_ptr.value
}

func (l *linkedlist) isPalindrome() bool {
    // 1. We get the list's middle
    slow_ptr := l.head
    fast_ptr := l.head

    for fast_ptr != nil && fast_ptr.next != nil {
        slow_ptr = slow_ptr.next
        fast_ptr = fast_ptr.next.next
    }
    // 2. Reverse second half from the middle
    head := slow_ptr
    var currentNode *node
    for head != nil {
        // We store next to swap it
        next := head.next
        // We reverse the next node
        head.next = currentNode
        // We go to next node
        currentNode = head
        head = next
    }

    // 3. We iterate the list while comparing it
    comparisonHead := l.head
    comparisonTail := currentNode
    for comparisonHead != nil && comparisonTail != nil {
        if comparisonHead.value != comparisonTail.value {
            return false
        }
        comparisonHead = comparisonHead.next
        comparisonTail = comparisonTail.next
    }

    return true
}

func main () {
    fmt.Println("Aoeu")

    l := linkedlist{}
    l.prependAll(1,2,3,2,1)

    l.deleteWithValue(0, true)

    l.printList()
    index, value := l.getMiddle()
    fmt.Printf("\nThe middle is: %d. Index: %d", value, index)

    /// Palindrome
    isPalindrome := l.isPalindrome()
    fmt.Printf("\nIs Palindrome: %v", isPalindrome)

    


} 
