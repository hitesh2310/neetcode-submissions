func containsNearbyDuplicate(nums []int, k int) bool {
    l := 0
    checkMap := make(map[int]int)
    for r:=0;r<len(nums);r++{
       if r -l > k {
            delete(checkMap,nums[l])
            l++
       }

       if _, exists := checkMap[nums[r]];exists {
        return true
       }
       checkMap[nums[r]]++
    }
return false
}
