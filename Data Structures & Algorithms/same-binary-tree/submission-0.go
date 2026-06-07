/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    res := dfs(p,q)

    return res 
}

func dfs(node1, node2 *TreeNode) bool {

    if node1 == nil && node2 == nil {
        return true
    }

    if (node1!=nil && node2==nil) || (node2!=nil && node1==nil){
        return false
    }
    if node1.Val !=node2.Val {
        return false
    }
    resL :=  dfs(node1.Left,node2.Left)
    resR := dfs(node1.Right,node2.Right)
    
    return (resL && resR)
}