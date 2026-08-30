func threeSum(nums []int) [][]int {
	res:= [][]int{}
	sort.Ints(nums)
	for i:=0;i<len(nums);i++ {
		if nums[i] > 0 {
			break
		}

		if i>0 && nums[i-1] == nums[i] {
			continue
		}

		l,r := i+1, len(nums)-1 
		for l<r {	
			if nums[i] + nums[l] + nums[r] == 0 {
				temp := []int{nums[i],nums[l],nums[r]}
				res = append(res,temp)
				l++
				r--
				    for l < r && nums[l] == nums[l-1] {
                    l++
                }

                for l < r && nums[r] == nums[r+1] {
                    r--
                }
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			}else{
				l++
			}
		}

	}
 return res 
}
