func maxSubArray(nums []int) int {
	sum := 0
	res := nums[0]
	for i:=0;i<len(nums);i++ {
		
		if sum < 0 {
			sum = 0	
		}
		sum = sum + nums[i]
		res = maxOf(res,sum)

	}
	return res
}



func maxOf(a,b int) int {
	if a > b {
		return a 
	}
	return b
}