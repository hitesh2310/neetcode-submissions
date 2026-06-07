func longestConsecutive(nums []int) int {
   MyMap := make(map[int]int)

     res := 0   
    for _,num:= range nums {
        MyMap[num] = 1 
    }
    for _,num := range nums {
        counter :=  longestForEach(num, MyMap)
        res= maxOf(res,counter)
    }
return res
}



func longestForEach(num int, MyMap map[int]int) int {
    counter := 1 
    currNum := num
    for {
        if _, exists:= MyMap[currNum-1]; !exists {
            return counter
        }else {
            counter++
        }
        currNum = currNum -1 
    }
}


func maxOf(a,b int) int {
    if a > b {
        return a
    }
    return b    
}