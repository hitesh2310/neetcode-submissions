/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
     if subRoot == nil {
        return true
     }
    subTreeString := ""
     subTreeString = serializeTree(subRoot,subTreeString)
     fmt.Println("String::", subTreeString)
     treeString := ""
     treeString = serializeTree(root, treeString)
     fmt.Println("String::", treeString)
     
     return strings.Contains(treeString, subTreeString)


    //  return true
}

func serializeTree(node *TreeNode,str string) string{
        if node == nil {
            return str+"X"
        }
        // fmt.Println("CurentNode::",node.Val)
        // fmt.Println("Current String::",str)
      
        str = str + strconv.Itoa(node.Val)
        // fmt.Println("Added curent node,  String::",str)
        str = serializeTree(node.Left,str)
        str = serializeTree(node.Right,str)
 
       return str
} 
    