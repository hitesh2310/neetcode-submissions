/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    counter :=  0
    res := math.MaxInt
    var  dfs func(node *TreeNode)   
    dfs = func(node *TreeNode) {
       
        if node == nil {
            return
        }
        fmt.Println("current element::", node.Val, "counter::",counter)
        dfs(node.Left)
        counter++
        if counter == k {
            res = node.Val
        }
        dfs(node.Right)
        
    }

    dfs(root)
    return res
}
