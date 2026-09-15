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
	visited := make(map[*Node]bool)
	var bfs func(node *Node)
	bfs = func(node *Node){

	 	q := []*Node{node}
		for len(q) > 0 {
			current := q[0]
			//fmt.Println("Checking value for ::", current)
			q = q[1:]

			if visited[current] {
				continue
			}

				if _, exists := oldToNew[current];!exists {
					newNode := &Node{
					Val: current.Val,
					Neighbors: []*Node{},
				}
				oldToNew[current] = newNode
				}
				
				for _,neighbor := range current.Neighbors {
						//fmt.Println("checking nerighbor", neighbor , " of current,", current)
						if _, exists := oldToNew[neighbor];!exists {
							//fmt.Println("new Node map to", neighbor, " not present")
							tempNode := &Node{
							Val: neighbor.Val,
							Neighbors: []*Node{},
							}
							oldToNew[neighbor] = tempNode
							//fmt.Println("Map oldToNew", oldToNew)
							//fmt.Println("Adding new clone of neighbor to clone of current", " neighbor::", tempNode, " clone of current::", oldToNew[current])
							oldToNew[current].Neighbors = append(oldToNew[current].Neighbors,tempNode)
						}else {
								//fmt.Println("Adding clone of neighbor to clone of current", " neighbor::", oldToNew[neighbor], " clone of current::", oldToNew[current])
							oldToNew[current].Neighbors = append(oldToNew[current].Neighbors,oldToNew[neighbor])
						}
					//fmt.Println("Marking current as visited true", current)
					visited[current] = true
					//fmt.Println("Adding neeighbor to queue", neighbor)
					q = append(q,neighbor)
				}
		} 			
	}

	bfs(node)
	return oldToNew[node]
}
