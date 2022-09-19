import("fmt")
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) (result bool) {
    values := make([]int, 0)
    
    // We create the int slice with the values
    for {
        values = append(values, head.Val)
        if head.Next == nil {
            break
        }
        
        head = head.Next
    }
    
    
    // We check if it's palindrome
    result = true
    for i,v := range values[:len(values)/2] {
        if values[ len(values) - i - 1] != v {
            result = false
            break
        }
        
    }
       
    return result
}

