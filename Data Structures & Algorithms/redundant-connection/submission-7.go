func findRedundantConnection(edges [][]int) []int {
	nodeMap := make(map[int][]int)
	visited := make(map[int]bool)
	var dfs func(node, parent int) bool 

	dfs = func(node, parent int) bool {
		if visited[node] {
			return false 
		}
		visited[node] = true
		for _, neighbor := range nodeMap[node] {
			if neighbor == parent {
				continue
			}
				if !dfs(neighbor, node){
					return false
				}
			}
		
		return true
	}

	for _,edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]],edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]],edge[0])
		for i:=1; i <=len(edges);i++ {
			visited[i] = false
		}
		if !dfs(edge[0],-1) {
			return []int{edge[0],edge[1]}
		}
	}

	return []int{}
}
