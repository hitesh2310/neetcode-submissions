/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {	
	var newHead ListNode
	cn1 := list1
	cn2 := list2

	for cn1!=nil && cn2!=nil {
		if cn1.Val < cn2.Val {
			AddNode(cn1.Val, &newHead)
			cn1 = cn1.Next 
		} else {
		 AddNode(cn2.Val, &newHead)
		 cn2 = cn2.Next     
		}
	}

	if cn1!=nil {
		for cn1 != nil {
			AddNode(cn1.Val, &newHead)
			cn1 = cn1.Next 
		}
	}


	if cn2!=nil {
		for cn2 != nil {
			AddNode(cn2.Val, &newHead)
			cn2 = cn2.Next 
		}
	}


res := newHead.Next
return res

}


func AddNode(x int, head *ListNode) {

	var tempNode ListNode
	tempNode.Val = x
	if head == nil {
		head = &tempNode
	}else {
		currNode := head 
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = &tempNode
	}
}