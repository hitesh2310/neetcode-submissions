func maxAreaOfIsland(grid [][]int) int {
    area := 0 
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	count := 0 

	var bfs func(i,j int)
	bfs = func(i, j int) {
	   count = 1
	    grid[i][j] = 0 
	   q := [][]int{{i,j}}
	   for len(q) >0 {
		 leftPop := q[0]
	   	 q = q[1:] 	
		 for _,dir := range directions {
			r, c := leftPop[0],leftPop[1]
			nr,nc := r + dir[0] , c + dir[1]
			if  nr <0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == 0 {
			continue
		    }
		    grid[nr][nc] = 0  //marking visited
		    //fmt.Println("Appending:", nr,", " ,nc, " pair")
			q = append(q, []int{nr,nc})
			//fmt.Println("q: ",q)
		    count++
			//fmt.Println("count after incr++:",count)
	     }
	   }
	}

	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == 1 {
			   //fmt.Println( "i=",i, "j=",j)
			   bfs(i,j)
			   //fmt.Println("Count=", count, " Area=", area)
			   area = maxOf(area,count)		
			}
		}
	}

	return area
}












func maxOf(a,b int) int {
	if a>b {
		return a
	}
	return b
} 
