func numIslands(grid [][]byte) int {
	count := 0 
	rows, cols := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	directions := [][]int{{-1,0},{1,0},{0,-1},{0,1}}

	var bfs func(r,c int)
	bfs = func(r,c int) {
		q:= [][]int{{r,c}}
		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			r,c = current[0],current[1]	
			if !visited[[2]int{r,c}] {
				fmt.Println("count++",r, c)
				count++
				visited[[2]int{r,c}] = true
				
			}
			for _, direction := range directions {
				nr,nc := r+direction[0], c+direction[1]

				if nr<0||nr >=rows || nc <0 || nc >=cols || grid[nr][nc]!='1' || visited[[2]int{nr,nc}]{
					continue
				}
				visited[[2]int{nr,nc}] = true
				q = append(q,[]int{nr,nc})
				}
		}
	}


	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == '1' {
			bfs(i,j) 
			}
		}
	}



	return count
}
