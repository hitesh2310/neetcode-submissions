func canFinish(numCourses int, prerequisites [][]int) bool {
    visited := make(map[int]bool)
	visiting := make (map[int]bool)
	courseMap := make(map[int][]int)

	for _, preReq:= range prerequisites {
		node, neighbor := preReq[0], preReq[1]
		courseMap[node] = append(courseMap[node], neighbor)

	}
	fmt.Println(courseMap)

	var dfs func(course int) bool 

	dfs = func(course int) bool {

		if visited[course] {
			return true
		}

		if len(courseMap[course]) == 0 {
			visited[course] = true
			return true
		}
		if visiting[course] {
			return false
		}

		visiting[course] = true

		for _, neighbor := range courseMap[course] {
			
			if !dfs(neighbor) {
				return false
			}
			visiting[course] = false
		}

		
		visited[course] = true
		return true
	}




	for course := range numCourses {
		if !dfs(course) {
			return false
		}

	}

	if len(visited)!= numCourses {
		return false
	}
	return true
}
