func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
	visiting := make(map[int]bool)
	visited := make(map[int]bool)
	
	preMap := make(map[int][]int)

	for _, preReq := range prerequisites {
		preMap[preReq[0]] = append(preMap[preReq[0]],preReq[1])
	}	

	fmt.Println(preMap)

	var dfs func(i int) bool 
	dfs = func(i int) bool {	

		if visiting[i] == true {
			return false
		}

		if visited[i]{
			return true
		}

		visiting[i] = true
		for _,preReq := range preMap[i] {
			if !dfs(preReq) {
				return false
			}
		}
		res = append(res,i)
	 	visited[i] = true
		visiting[i] = false
		return true
	}




	for numCrs := range numCourses {
		if !dfs(numCrs) {
			return []int{}
		}
	}

	return res 
}
