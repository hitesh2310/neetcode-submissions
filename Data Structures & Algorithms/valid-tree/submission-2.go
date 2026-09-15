func validTree(n int, edges [][]int) bool {
    nodeMap := make(map[int][]int)
	visited := make(map[int]bool)

	for _,edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]], edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]], edge[0])
	}
	fmt.Println(nodeMap)

	var dfs func(i int, parent int ) bool 
	dfs = func(i int, parent int ) bool {

		if visited[i] {
			return false
		} 

		visited[i] = true
		for _, node := range nodeMap[i] {
			
			if node == parent {
				continue
			}
			
			if !dfs(node,i) {
				return false
			}
		}
		return true
	}

	if !dfs(0,-1) {
		return false
	}


if len(visited) != n {
    return false
}
	return true
}
