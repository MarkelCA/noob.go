import("fmt")
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) (result bool) {
    originalHead := head
    // 1. We get the list's middle
    slow_ptr := head
    fast_ptr := head

    for fast_ptr != nil && fast_ptr.Next != nil {
        slow_ptr = slow_ptr.Next
        fast_ptr = fast_ptr.Next.Next
    }
    // 2. Reverse second half from the middle
    reversedHead := slow_ptr
    var currentNode *ListNode
    for reversedHead != nil {
        // We store Next to swap it
        next := reversedHead.Next
        // We reverse the next node
        reversedHead.Next = currentNode
        // We go to next node
        currentNode = reversedHead
        reversedHead = next
    }

    // 3. We iterate the list while comparing it
    comparisonHead := originalHead
    comparisonTail := currentNode
    for comparisonHead != nil && comparisonTail != nil {
        if comparisonHead.Val != comparisonTail.Val {
            return false
        }
        comparisonHead = comparisonHead.Next
        comparisonTail = comparisonTail.Next
    }

    return true

}
