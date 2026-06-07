/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    first := l1
    second := l2
    res:= make([]int,0)
    tensPlace := 0
    for first!=nil || second!=nil {
        firstNumber := 0
        secondNumber := 0
        if first != nil {
            firstNumber = first.Val
        }
         if second != nil {
            secondNumber = second.Val
        }
        temp := firstNumber + secondNumber + tensPlace 
        if temp > 9 {
            //split 
            unitPlace := temp % 10
            tensPlace = temp/10 
            res = append(res,unitPlace)
        }else {
             res = append(res,temp)
             tensPlace = 0
        }

        if first!=nil {
        first = first.Next
        }
        if second!=nil {
        second = second.Next
        }
    }
    if tensPlace > 0 {
        res = append(res,tensPlace)
    }
    fmt.Println(res)
    var Reshead *ListNode
    var Current *ListNode

    for _,num := range res {
            tempNode :=  &ListNode{Val:num}
            tempNode.Val = num
            tempNode.Next = nil
        
        if Reshead == nil {
            Reshead = tempNode
        }else{
            Current.Next = tempNode
        }
        Current = tempNode
    } 
    return Reshead
}