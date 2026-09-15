func orangesRotting(grid [][]int) int {
    rows := len(grid)
	cols:= len(grid[0])
	directions := [][]int{{-1,0},{1,0},{0,1},{0,-1}}
	steps:= -1
	visited := make(map[[2]int]bool)
	q := [][]int{}

	count := 0
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == 2 {
				q = append(q,[]int{i,j})
				visited[[2]int{i,j}] = true
				count++
			}
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	if count == 0 {
		return 0
	}


	for len(q) > 0 {

		size := len(q)

		for k:=0;k<size;k++ {
			r,c := q[0][0], q[0][1]
			q = q[1:]

			for _,dir := range directions {
				nr,nc := r+dir[0], c+dir[1]

				if nr<0 || nr >=rows || nc < 0 || nc >= cols || grid[nr][nc] == 0 || visited[[2]int{nr,nc}] {
					continue 
				}
				
				q = append(q,[]int{nr,nc})
				visited[[2]int{nr,nc}] = true
			}

		}
		steps ++
	}

	if count != len(visited) {
		return -1 
	}
	return steps
}
