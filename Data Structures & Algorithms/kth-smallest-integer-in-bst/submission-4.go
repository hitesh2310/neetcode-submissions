/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Element struct {
	Val int
}


type MinHeap  []*Element

/*
1. Less
2.Len
3. swap 
4. push 
5. pop
*/

func (h MinHeap) Less(i,j int) bool {
	return h[i].Val<h[j].Val
}

func (h MinHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h,x.(*Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(root *TreeNode, k int) int {
    var res any
	pq := &MinHeap{}
	heap.Init(pq)
	var pot func(node *TreeNode) 
	pot = func(node *TreeNode) {

		if node == nil {
			return 
		}
		// push element in 
		element:= &Element{ Val: node.Val}
		heap.Push(pq,element)
		if node.Left != nil {
		pot(node.Left)
		}
		if node.Right != nil {
			pot(node.Right)
		}
	}

	pot(root)
	// fmt.Println(pq)
	i:=1
	for i < k {
		heap.Pop(pq)
		i++
	}
	res = heap.Pop(pq)
	return res.(*Element).Val 
}
