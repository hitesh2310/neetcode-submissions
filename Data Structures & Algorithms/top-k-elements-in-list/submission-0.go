func topKFrequent(nums []int, k int) []int {

    bucket := make([][]int,len(nums)+1)
    myMap := make(map[int]int)
    counter :=0
    res := make([]int,0)
    for i:=0;i<len(nums);i++{
        myMap[nums[i]] = myMap[nums[i]]+1
    }

    // fmt.Println("Map::", myMap)

    for key,val := range myMap {
        var temp []int 
        temp = append( bucket[val],key)
        bucket[val] = temp
    }

    fmt.Println("Bucket::", bucket, len(bucket))
    
    for i:=len(bucket)-1;i>0;i--{
        fmt.Println("i::",i, bucket[i])
        
            for _, val := range bucket[i] {
                res = append(res,val)
                counter++
                if counter == k {
                    return res
                }
            }    
    }
    return res
}
