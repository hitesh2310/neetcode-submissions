/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0 

	var dfs func(node *TreeNode) int 
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		result = maxOf(result, left + right)
		return 1 + maxOf(left, right)
	}

	dfs(root)
	return result 
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

