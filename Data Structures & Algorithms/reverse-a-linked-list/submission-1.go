/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    
    iterator := head 
    var prev *ListNode
    for iterator!=nil {
        var tempNode *ListNode
        tempNode = iterator.Next
        iterator.Next = prev
        prev = iterator
        iterator = tempNode
    }

    return prev 
}
