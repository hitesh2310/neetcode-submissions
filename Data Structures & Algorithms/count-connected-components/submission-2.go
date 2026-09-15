func countComponents(n int, edges [][]int) int {
    graphs := 0
	nodeMap := make(map[int][]int)
	visited := make(map[int]bool)
	for _, edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]], edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]], edge[0])
	}

	var dfs func(node, parent int) 
	dfs = func (node,parent int) {
		visited[node] = true
		for _, neighbor := range nodeMap[node] {
			if neighbor == parent {
				continue
			}
			if !visited[neighbor] {
			dfs(neighbor,node)	
			}
		}
	}
	
	for i := 0; i < n; i++ {
    if !visited[i] {
        dfs(i, -1)
        graphs++
    }
}

	return graphs
}
