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

	var prevNode *ListNode
	currNode := head 
	// fmt.Println("Curr = Head::", currNode.Val)
	for currNode != nil {
	// 	fmt.Println("currNode", currNode.Val)		
	// 	fmt.Println("currNode.Next", currNode.Next.Val)		
		tempNode := currNode.Next
		// fmt.Println("tempNode::", tempNode.Val)
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode

	}
	return prevNode
}