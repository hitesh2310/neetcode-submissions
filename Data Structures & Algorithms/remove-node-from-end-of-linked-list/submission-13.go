/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //1. set up leading ptr 
    counter := 0 
    leadPtr := head
    for  counter<n {
        leadPtr = leadPtr.Next
        counter++
    }

    // curr:=head
    dummy := &ListNode{Next:head}
    curr := dummy
    for leadPtr!=nil {
      leadPtr = leadPtr.Next
      curr = curr.Next
    }


    curr.Next = curr.Next.Next
        
    return dummy.Next
}
