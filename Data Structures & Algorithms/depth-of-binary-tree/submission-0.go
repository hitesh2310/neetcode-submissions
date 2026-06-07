/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    
   if root == nil {
    return 0
   }

    return 1+ maxOf(maxDepth(root.Left),maxDepth(root.Right))
}

func maxOf(a,b int) int {
    if a > b {
        return a 
    }
    return b
}