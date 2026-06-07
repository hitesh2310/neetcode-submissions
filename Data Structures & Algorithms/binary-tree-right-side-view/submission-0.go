/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    bfsResult := [][]int{}
    res := []int{}
    var dfs func(node *TreeNode,depth int) 
    dfs = func(node *TreeNode,depth int) {
        if node == nil {
            // bfsResult[depth] = append(bfsResult[depth],-1)
            return 
        }
    
        if len(bfsResult) == depth {
            bfsResult = append(bfsResult,[]int{})
        }
          
        bfsResult[depth] = append(bfsResult[depth],node.Val)
        dfs(node.Left, depth+1)
        dfs(node.Right, depth+1)        
    }
    dfs(root,0)
    fmt.Println(bfsResult)

    for _,eachArray := range bfsResult {

        res = append(res,eachArray[len(eachArray)-1])

    }
    return res
}