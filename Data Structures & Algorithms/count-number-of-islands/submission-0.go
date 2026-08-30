func numIslands(grid [][]byte) int {
	directions := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	numberOfIslands := 0
	// isVisited := make(map[[]int]bool)   // Set to maintain visited "1" nodes; not needed 
	rows:= len(grid)
	cols:= len(grid[0])

	var bfs func(i,j int)
	bfs = func(i,j int) {
		q:= [][]int{{i,j}}
		grid[i][j] = '0'   //setting value as 0(0/1) ; marking this as visited 

		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			row,col := front[0], front[1]
			for _, dir := range directions {
				nr, nc := row+dir[0], col + dir[1]
				if nr < 0 || nr >= rows || 
				 nc < 0 || nc >= cols || grid[nr][nc] == '0' {
					continue
				 }
				 q = append(q, []int{nr, nc})
				 grid[nr][nc] = '0'
			}
		}
	}

	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == '1' {
				bfs(i,j)
				numberOfIslands++
			}
		}
	}




	return numberOfIslands	
}
