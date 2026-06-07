/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {


    if head.Next==nil{
        return false
    }
    oneStepIterator, twoStepIterator := head, head.Next


    for oneStepIterator!=nil && twoStepIterator!=nil {

        if oneStepIterator == twoStepIterator  {
            return true
        }
        oneStepIterator = oneStepIterator.Next
        if twoStepIterator.Next==nil { 
            return false
        }

        twoStepIterator = twoStepIterator.Next.Next
        
    }
    


    return false
    
}
