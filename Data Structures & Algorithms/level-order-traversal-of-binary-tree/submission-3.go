/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

    res := [][]int{}

	var bfs func(node *TreeNode)
	bfs = func(node *TreeNode) {
		
		q:=[]*TreeNode{node}
		
		for len(q) > 0 {
			size := len(q)
			temp := []int{}
			for k:=0;k<size;k++ {
				current := q[0]
				q = q[1:]
				temp = append(temp,current.Val)
				if current.Left !=nil {
					q = append(q,current.Left)
				} 
				if current.Right!=nil {  
				  q = append(q,current.Right)				
				}				
			}
			res = append(res,temp)
		}
	} 


	bfs(root)
	return res
}
