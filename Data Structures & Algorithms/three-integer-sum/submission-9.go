func threeSum(nums []int) [][]int {
res:= [][]int{}
sort.Ints(nums)
l,r := 0,len(nums)-1

for i:=0;i<len(nums);i++ {

    a := nums[i]
    if a>0 {
        break
    }
    if i>0 && nums[i]==nums[i-1] {
        continue
    }

    l = i+1
    r = len(nums)-1
    for l < r {
        sum := nums[i]+nums[l]+nums[r]

        if sum < 0 {
            l++
        }else if sum > 0 {
            r -- 
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
