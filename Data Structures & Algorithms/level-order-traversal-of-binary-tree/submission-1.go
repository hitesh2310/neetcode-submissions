/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    bfsResult := [][]int{}

    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int){

        if node == nil {
            return 
        }

        if level == len(bfsResult) {
            temp := []int{}
            bfsResult = append(bfsResult,temp)
        }
        dfs(node.Left,level+1)
        bfsResult[level] = append(bfsResult[level],node.Val)    
        dfs(node.Right,level+1)
    }


    dfs(root,0)
    return bfsResult
}
