func findRedundantConnection(edges [][]int) []int {
  //  var result [][]int
	nodeMap := make(map[int][]int)
	visited := make(map[int]bool)

	
	
	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		//fmt.Println("node::",node,"parent::",parent)
		if visited[node] {
			return false
		}	
		//fmt.Println("Neighbors::",nodeMap[node])
		//fmt.Println("Marking visited for node::", node)
		visited[node] = true
		for _, neighbor := range nodeMap[node] {
			if neighbor == parent {
				continue
			}
			if !dfs(neighbor,node){
				return false
			}
		}
		return true
	}


	n:=len(edges)
	for _,edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]],edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]],edge[0])		
		//fmt.Println(nodeMap)
  for i := 0; i <= n; i++ {
            visited[i] = false
        }
		notCycle := dfs(edge[0],-1)
	
		//fmt.Println("notCycle", notCycle)
		if !notCycle {
			//result = append(result,[]int{edge[0],edge[1]})
			return []int{edge[0],edge[1]}
		}
		//fmt.Println(result)
	}
	


	
	return []int{}
}
