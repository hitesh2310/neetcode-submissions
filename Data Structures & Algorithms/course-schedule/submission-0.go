func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make(map[int]bool)
	preReqMap := make(map[int][]int)

	for _, preReq := range prerequisites {
		preReqMap[preReq[0]] = append(preReqMap[preReq[0]],preReq[1])
	}

	fmt.Println(preReqMap)

	var dfs func(i int) bool 
	dfs = func(i int) bool {
		

		if visited[i] {
			return false
		}
		if len(preReqMap[i])==0 {
			return true
		}

		visited[i] = true 
		for _,preReq := range preReqMap[i] {
			if !dfs(preReq) {
				return false
			}
		}
		visited[i] = false
		preReqMap[i] = []int{}

		return true
	}


	for i:=0;i<numCourses;i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
