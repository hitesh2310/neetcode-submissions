func countComponents(n int, edges [][]int) int {
	count := 0 
	visited := make(map[int]bool)
	nodeMap := make(map[int][]int)

	for _,edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]],edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]],edge[0])
	}

	var dfs func(node, parent int) 
	dfs = func(node, parent int) {
			if visited[node] {
				return 
			}
			visited[node] = true
			for _, neighbor := range nodeMap[node] {
				if neighbor == parent  {
					continue
				}
				dfs(neighbor,node)
			}
		
	}


	for _ , edge := range edges {
		if !visited[edge[0]] {
			fmt.Println("count ++", edge[0])
			count ++ 
			dfs(edge[0],-1)
		}
	}

	if len(visited) == n {
		return count
	}
	return count + n-len(visited)
}
