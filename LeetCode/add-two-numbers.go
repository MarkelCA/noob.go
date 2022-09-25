package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

type linkedlist struct {
    head *ListNode
}

func main() {
    fmt.Println("e")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ll := linkedlist{}

    current1 := l1
    current2 := l2
    
    plusOne := false
    for current1 != nil && l2 != nil {
        
        var val1 int
        var val2 int
        
        if current1 != nil {
            val1 = current1.Val
        } else {
            val1 = 0
        }
        
        if current2 != nil {
            val2 = current2.Val
        } else {
            val2 = 0
        }
        
        
        sum := val1 + val2
        
        if plusOne {
            sum ++
        }
        
        if sum >= 10 {
            sum -= 10
            plusOne = true
        } else {
            plusOne = false
        }
        
        ll.appendEnd(sum)
        
        current1 = current1.Next
        current2 = current2.Next
    }
    
    return ll.head
}

func (l *linkedlist) appendEnd(value int) {
    if l.head == nil {
        l.head = &ListNode{Val:value}
    } else {
        currentNode := l.head
        for currentNode.Next != nil {
            currentNode =  currentNode.Next
        }

        currentNode.Next = &ListNode{Val:value}
    }
}

