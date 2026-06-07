func hasDuplicate(nums []int) bool {

   myMap := make(map[int]int)

   for i, num := range nums {

    if _, exists := myMap[num];exists{
        return true
    }
    myMap[num] = i 
   }

   return false 
}
