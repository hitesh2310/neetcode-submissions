func twoSum(nums []int, target int) []int {
        myMap := make(map[int]int)
        for i,num := range nums {
            diff := target - num 
            if val ,exists := myMap[diff]; exists {
                return []int{val, i}
            }
            myMap[num] = i     
        }

    return []int{-1,-1}
}
