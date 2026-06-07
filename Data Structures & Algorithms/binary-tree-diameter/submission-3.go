/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0 
    dfs(root, &diameter)
    return diameter
}

func dfs(root *TreeNode, diameter *int) int {

    if root == nil {
        return 0
    }

    left := dfs(root.Left,diameter)
    right := dfs(root.Right,diameter)

    *diameter = maxOf(*diameter,left + right)

    return 1 + maxOf(left,right)
}


func maxOf(a,b int) int  {
    if a >b {
        return a
    }
    return b
}