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

	if head.Next == nil {
		return head
	}
		currNode := head 

		var prevNode *ListNode
		for currNode != nil {

			tempNode := currNode.Next
			currNode.Next = prevNode
			prevNode=currNode
			currNode = tempNode
		}

	return prevNode
}
