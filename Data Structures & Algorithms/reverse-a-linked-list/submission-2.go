/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    
    if head == nil {
        return nil
    }
    current := head 
    var prev *ListNode 
    for current!=nil {
        temp:= &ListNode{}
        temp = current.Next
        current.Next = prev
        prev = current
        current = temp
    }

    return prev

}
