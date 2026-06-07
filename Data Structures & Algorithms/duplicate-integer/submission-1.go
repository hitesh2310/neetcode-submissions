func hasDuplicate(nums []int) bool {
 
 myMap := make(map[int]int)
 for _, num := range nums {
    if _, exists := myMap[num];exists {
        return true
    }
    myMap[num] = 1
 }

return false
}
