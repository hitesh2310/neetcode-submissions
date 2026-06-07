/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
      if head == nil || head.Next == nil {
        return
    }

    oneStepPtr := head
    twoStepPtr := head.Next
//  1. Find Mid 
    for twoStepPtr != nil && twoStepPtr.Next!=nil  {
        twoStepPtr = twoStepPtr.Next.Next
        oneStepPtr = oneStepPtr.Next
    } 

    // fmt.Println("Mid::",oneStepPtr.Val)
//  2. Reverse the second list, second list starts from mid + 1.
    second := oneStepPtr.Next 
    oneStepPtr.Next = nil 
    var prev *ListNode 
    for second != nil {
        tempNode := second.Next 
        second.Next = prev 
        prev = second
        second = tempNode
    }  
    fmt.Println("head of reveres list::", prev.Val)
//  3. Merge the firs and second list 
    first := head 
    second = prev

    for second != nil {
        tmp1 := first.Next
        tmp2 := second.Next
        first.Next = second 
        second.Next = tmp1
        first = tmp1 
        second = tmp2
    }
}
