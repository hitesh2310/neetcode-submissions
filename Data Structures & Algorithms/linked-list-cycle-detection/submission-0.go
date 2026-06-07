/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

   if head == nil {
        return false
    }
    oneStepPtr := head
    twoStepPtr := oneStepPtr.Next 
    
    
    for oneStepPtr != nil && twoStepPtr != nil {
        
        if oneStepPtr == twoStepPtr {
            return true
        }
        oneStepPtr = oneStepPtr.Next
        if twoStepPtr.Next == nil {
            return false
        }
        twoStepPtr = twoStepPtr.Next.Next
    }    


    return false
}
