/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    
    _,resB := dfs(root)

    return resB
}


func dfs(node *TreeNode) (int,bool) {
    if node == nil {
        return 0,true
    }
    left, Lbalanced  := dfs(node.Left)
    right, Hbalanced := dfs(node.Right)
    if !Lbalanced || !Hbalanced || ((right-left) >=2 || (right - left)<=-2) {
        return -1, false
    } 

    return 1+maxOf(left,right), true   
}

func maxOf(a,b int) int {
    if a > b {
        return a
    }
    return b
}