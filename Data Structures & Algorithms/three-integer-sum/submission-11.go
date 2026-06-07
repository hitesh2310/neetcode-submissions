func threeSum(nums []int) [][]int {
res:= [][]int{}
sort.Slice(nums, func(a,b int) bool {
    return nums[a] < nums[b]
})

fmt.Println(nums)


for i:=0;i<len(nums);i++ {
    if nums[i]> 0  {
       break
    }
    if i > 0 && nums[i] == nums[i-1] {
        continue
    }
    l,r := i+1, len(nums)-1
    for l < r {
        sum := nums[i] + nums[l] + nums[r]

        if sum < 0 {
            l++
        }else if sum > 0 {
            r--
        }else {
            tempRes := []int{nums[i],nums[l],nums[r]}
            res = append(res,tempRes)
            l++
            r--
             for l < r && nums[l] == nums[l-1] {
                    l++
                }
        }
    }

}

return res
}
