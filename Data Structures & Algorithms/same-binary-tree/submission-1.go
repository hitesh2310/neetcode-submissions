/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {


	var pOT func(nodeP *TreeNode, nodeQ *TreeNode) bool
	pOT = func(nodeP *TreeNode, nodeQ *TreeNode) bool {

		if nodeP == nil && nodeQ!=nil || nodeQ==nil && nodeP!=nil {
			fmt.Println("p/q is null")
			return false
		}

		if nodeP == nil && nodeQ == nil {
			return true
		}

		//check val of p and q 
		fmt.Println(p.Val, " <->", q.Val)
		if nodeP.Val != nodeQ.Val {
			return false
		}

		
		
		return pOT(nodeP.Left,nodeQ.Left) && pOT(nodeP.Right,nodeQ.Right)
	}


	

 return pOT(p,q)    
}
