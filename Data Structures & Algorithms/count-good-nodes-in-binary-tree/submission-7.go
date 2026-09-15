/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
	maxSoFar := -101	 
	var pot func(node *TreeNode, maxSoFar int) 
	pot = func(node *TreeNode, maxSoFar int) {

		if node == nil {
			return
		}

		// check maxSoFar < node 
		if maxSoFar <= node.Val {
			// fmt.Println("Found::", node.Val)
			count ++ 
		}
		maxSoFar = MaxOf(node.Val, maxSoFar)

		if node.Left != nil {
			pot(node.Left,maxSoFar)
		}

		if node.Right != nil {
			pot(node.Right, maxSoFar)
		}

	}

	
	pot(root, maxSoFar)
	return count
}




func MaxOf(a, b int) int {
	if a > b {
		return a 
	}
	return b
}