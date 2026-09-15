/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    //sol: find the node from where p<=root && q>root 
	res:= root 


	var pOT func(node *TreeNode, p *TreeNode, q *TreeNode) 
	pOT = func(node *TreeNode, p *TreeNode, q *TreeNode)  {
		if node == nil {
			return 
		}

		if (p.Val <= node.Val && q.Val >= node.Val)  {
			res = node
			return
		}

		if (q.Val <= node.Val && p.Val >= node.Val) {
			res = node
			return
		}
		
		pOT(node.Left,p,q)
		pOT(node.Right,p,q)
	}

	pOT(root,p,q)
	return res
}
