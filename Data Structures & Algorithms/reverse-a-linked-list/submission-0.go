/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    currentNode := head
    var prev *ListNode
    for currentNode != nil {
        fmt.Println(currentNode.Val)
        tempNode := currentNode.Next
        currentNode.Next = prev 
        prev = currentNode
        currentNode = tempNode
    } 
    fmt.Println("CurrentNode::", prev)
return prev
}
