func productExceptSelf(nums []int) []int {
res:=[]int{}
productMap := make(map[int]int)
product := 1
for i:=len(nums)-1;i>=0;i--{
    if i == len(nums)-1 {
        productMap[i]=product
    }else{
        product = product*nums[i+1]
        fmt.Println(i, product)
        productMap[i]=product
    }
}
fmt.Println(productMap)
product = 1 
for i:=0;i<len(nums);i++{
    res = append(res,productMap[i]*product)
    product = product * nums[i]
}

return res
}
