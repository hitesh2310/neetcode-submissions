func combinationSum2(candidates []int, target int) [][]int {
    
    sort.Slice(candidates,func(a,b int)bool {
        return candidates[a] < candidates[b]
    })
    fmt.Println(candidates)
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
            return
        }
        
        if i>=len(candidates) || sum > target{
            return 
        }


        subset = append(subset,candidates[i])
        backtrack(i+1)
        for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
            i++
        }
        subset = subset[:len(subset)-1]
        backtrack(i+1)
    }
    
    backtrack(0)
    return res 
}
