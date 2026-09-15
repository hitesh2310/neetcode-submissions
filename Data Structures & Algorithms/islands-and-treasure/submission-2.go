func islandsAndTreasure(grid [][]int) {
	rows := len(grid)
	cols:= len(grid[0])
	directions := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	q:= [][]int{}
	visited := make(map[[2]int]bool)

	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == 0 {
					q = append(q,[]int{i,j})
					visited[[2]int{i,j}] = true 
			}
		}
	}

	dist := 0
	
	for len(q) > 0 {
	
		levelSize := len(q)
		for k:=0;k<levelSize;k++ {
		r, c := q[0][0],q[0][1]
		grid[r][c] = dist
		q = q[1:]
		
		for _, dir := range directions {
			nr := r + dir[0]
			nc := c + dir[1]

			if nr < 0 || nr >= rows || nc <0 || nc >= cols || 
			grid[nr][nc]==-1 || visited[[2]int{nr,nc}]{
				continue
			} 
			
			q = append(q,[]int{nr,nc})
			visited[[2]int{nr,nc}] = true
		  }
		}
		dist++
	}
}
