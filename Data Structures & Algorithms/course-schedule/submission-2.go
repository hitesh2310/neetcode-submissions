func canFinish(numCourses int, prerequisites [][]int) bool {
	visiting := make(map[int]bool)
	preReqMap := make(map[int][]int)
	visited := make(map[int]bool)

	for _, preReq := range prerequisites {
		preReqMap[preReq[0]] = append(preReqMap[preReq[0]],preReq[1])
	}

	fmt.Println(preReqMap)

	var dfs func(i int) bool 
	dfs = func(i int) bool {
		

		if visiting[i] {
			return false
		}
		if visited[i] {
			return true
		}

		visiting[i] = true 
		for _,preReq := range preReqMap[i] {
			if !dfs(preReq) {
				return false
			}
		}
		visiting[i] = false
		visited[i] = true

		return true
	}


	for i:=0;i<numCourses;i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
