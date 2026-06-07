func maxArea(nums []int) int {

    l,r := 0, len(nums)-1
    max := 0
    for l<r {
            area:= (r-l) *minOf(nums[l],nums[r])
            if area > max {
                max = area
            }

            if nums[l]> nums[r] {
                r--
            }else {
                l++
            }
            

    }

    return max
}


func minOf(a,b int) int {

    if a < b {
        return a
    }

    return b 
}