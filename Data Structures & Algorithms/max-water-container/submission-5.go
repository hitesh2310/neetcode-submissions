func maxArea(heights []int) int {

	l,r :=0, len(heights) -1 
	res:=0
	for l <r {
		currRes := minOf(heights[l], heights[r]) * (r - l) 
		res = maxOf(currRes,res)

		if heights[l] > heights[r]{
			r--
		}else{
			l++
		}

	}

return res 
}




func minOf(a,b int) int {
	if a < b {
		return a
	}else {
		return b 
	}
}


func maxOf(a,b int) int {
	if a > b {
		return a 
	}else {
		return b
	}
}