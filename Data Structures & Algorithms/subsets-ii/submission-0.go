func subsetsWithDup(nums []int) [][]int {

    sort.Ints(nums)
    fmt.Println(nums)
    res:= [][]int{}
    var backtrack func(i int)
    subset := []int{}

    backtrack =  func(i int){

        if i>=len(nums){
            temp := make([]int,len(subset))
            copy(temp,subset)
            res = append(res,temp)
            return
        }
       
        subset = append(subset,nums[i])
        backtrack(i+1)
        
        for i+1<len(nums) && nums[i+1] == nums[i] {
            i++
        }
        subset = subset[:len(subset)-1]
        backtrack(i+1)

    } 

    backtrack(0)
    return res
}
