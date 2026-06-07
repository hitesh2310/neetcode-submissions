func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)

    for index, num := range nums{
        diff := target - nums[index]
        if val,exists := indexMap[diff]; exists {
            return []int{val,index}
        }
        indexMap[num] = index
    }

    return []int{}
}
