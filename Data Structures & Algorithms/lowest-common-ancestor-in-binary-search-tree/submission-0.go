/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
     
     current := root 

     for current!=nil {

        if current.Val < p.Val && current.Val < q.Val {
            current = current.Right
        }else if current.Val > p.Val && current.Val > q.Val {
            current = current.Left
        }else  {
            return current
        }
     }

     return nil
}

