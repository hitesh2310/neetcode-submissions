/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)
	oldToNew[node] = &Node{
		Val : node.Val,
		Neighbors: []*Node{} ,
	}

	q := []*Node{node}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]


		for _,neighbor := range current.Neighbors {
			if _,exists := oldToNew[neighbor];!exists {
				oldToNew[neighbor] = &Node{
					Val : neighbor.Val,
					Neighbors: []*Node{},
				}
			q  = append(q,neighbor)
		}
		oldToNew[current].Neighbors = append(oldToNew[current].Neighbors,oldToNew[neighbor])
		}
	}

	return oldToNew[node]
}
