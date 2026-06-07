/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    first := head 
    counter := 0 
    for counter < n {
        first = first.Next 
        counter++
    }

    dummy := &ListNode{Next:head}
    second := dummy
    for first!=nil {
        first = first.Next
        second = second.Next
    }

    second.Next = second.Next.Next

    return dummy.Next
}
