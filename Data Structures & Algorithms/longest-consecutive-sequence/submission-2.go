func longestConsecutive(nums []int) int {
myMap := make(map[int]int)
res := 0

for _ , num := range nums{
    myMap[num]=1
}

for _,num := range nums {
    fmt.Println("checking::", num)
    currRes := 1 
    curr := num 
    for {
        if _, exists := myMap[curr-1];exists {
            currRes++ 
            curr = curr - 1 
        }else{
            if currRes > res {
                res = currRes   
            }
            break
        }
        
    }
}

return res
}
