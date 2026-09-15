/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := [][]int {}
	var q []*TreeNode

	var bfs func() 
	bfs = func() {
		for len(q) > 0 {
		size := len(q)
		levelRes := []int{}
		for k:=0;k<size;k++ {
			current := q[0]
			q = q[1:]
			if current != nil {
				levelRes = append(levelRes,current.Val)
			}

			if current.Left != nil {
				q = append(q,current.Left)
			}

				if current.Right != nil {
				q = append(q,current.Right)
				}
			
			}
		res = append(res,levelRes)
		}

	}
	
	if root == nil {
		return res
	}
	q  = []*TreeNode{root}
	bfs()
	return res 
}
