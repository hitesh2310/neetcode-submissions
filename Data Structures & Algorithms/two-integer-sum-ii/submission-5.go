func twoSum(numbers []int, target int) []int {
res:= []int{}

l,r:=0,len(numbers)-1

for l<r {

    if numbers[l] + numbers[r] == target {
        return []int{l+1,r+1}
    }else if numbers[l] + numbers[r] > target {
        r --
    }else {
        l++
    }
}

return res
}
