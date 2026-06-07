func twoSum(nums []int, target int) []int {

 myMap := make(map[int]int)


 for i:=0;i<len(nums);i++ {

    diff := target - nums[i] 
    if val, exists := myMap[diff] ; exists {
        return []int{val,i}
    }
    
    myMap[nums[i]] = i

 }  

 return []int{0,0} 
}
