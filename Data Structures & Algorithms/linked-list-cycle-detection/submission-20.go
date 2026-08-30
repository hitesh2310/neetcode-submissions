/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	tempMap := make(map[interface{}] interface{})	
	currNode := head 

	for currNode != nil {

		if _, exists := tempMap[currNode] ; exists {
			return true
		}else {
			tempMap[currNode] = currNode 
		}

		currNode = currNode.Next
	}

	return false
}
