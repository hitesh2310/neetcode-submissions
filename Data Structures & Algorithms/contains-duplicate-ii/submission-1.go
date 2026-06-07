func containsNearbyDuplicate(nums []int, k int) bool {


    l := 0
    checkMap := make(map[int]int)
    counter :=0
    for r:=0;r<len(nums);r++{
        if _, exists := checkMap[nums[r]];exists {
            return true
        } else {
            checkMap[nums[r]] = checkMap[nums[r]]+1
        }
        if counter == k {
            l++
            r=l
            counter=0
            checkMap  =  make(map[int]int)
        }else{
            counter++
        }
        
    }

return false
}
