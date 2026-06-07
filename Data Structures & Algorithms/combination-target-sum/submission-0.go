func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}

    var backtrack func(i int) 
    subset := []int{}
    backtrack = func(i int) {
        fmt.Println("current subset::", subset)
        sum :=0
        for _,each := range subset {
            sum = sum + each
        }

        if sum == target {
            fmt.Println("----")
            fmt.Println(subset)
            fmt.Println(sum)
            temp := make([]int,len(subset))
            copy(temp,subset)
            res = append(res,temp)
        }
        
        if i>=len(nums) || sum>=target {
            return 
        }

        subset = append(subset,nums[i])
        backtrack(i)
        subset = subset[:len(subset)-1]
        backtrack(i+1)
    }
    
    backtrack(0)
    return res 
}
