/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
        
        if head==nil || head.Next == nil {
            return false
        }
        oneHopper := head
        twoHopper:= head.Next

        for oneHopper!=nil && twoHopper!=nil {

            if oneHopper == twoHopper {
                return true
            }

            oneHopper = oneHopper.Next 
            if twoHopper.Next==nil {
                return false
            }
           
            twoHopper = twoHopper.Next.Next
            
        }

    return false
}
