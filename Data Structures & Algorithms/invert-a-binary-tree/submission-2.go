/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
        inorderTraversal(root)        


        return root

}

func inorderTraversal(node *TreeNode) {
    if node != nil {

        if node.Left!=nil {
            inorderTraversal(node.Left)
        }
        
        if node.Right!=nil {
            inorderTraversal(node.Right)
        }
        swapNode(node)
    }
    return 
}

func swapNode(node *TreeNode) {
    if node == nil {
        return
    }
    
        node.Left, node.Right = node.Right,node.Left
    
}
