func validTree(n int, edges [][]int) bool {
    visited := make(map[int]bool)

	nodeMap := make(map[int][]int)

	for _,edge := range edges {
		parent, node := edge[0],edge[1]
		nodeMap[parent] = append(nodeMap[parent],node)
		nodeMap[node] = append(nodeMap[node],parent)
	}
	fmt.Println(nodeMap)

	var dfs func(x, parent int) bool 
	dfs = func(x, parent int) bool {
		if visited[x] {
			return false
		}

		visited[x] = true
		for _ , neighbor := range nodeMap[x] {

			if parent == neighbor {
				continue
			}

			if !dfs(neighbor, x) {
				return false
			}
		}
		return true
	}
	


	if !dfs(0,-1){
		return false
	}

	if  len(visited) != n {
		return false
	}

	return true
}
