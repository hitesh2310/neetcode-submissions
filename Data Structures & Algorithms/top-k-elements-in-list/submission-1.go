func topKFrequent(nums []int, k int) []int {
res := []int{}
bucket := make([][]int,len(nums)+1)
map1:= make(map[int]int)
for _ ,num := range nums {
    map1[num]++
}

for key,val := range map1 {
    bucket[val] = append( bucket[val],key) 
}

fmt.Println(bucket)
for i:=len(bucket)-1;i>0;i-- {
    res = append(res, bucket[i] ...)
    if len(res) >=k {
        return res
    }   
}


return res
}
