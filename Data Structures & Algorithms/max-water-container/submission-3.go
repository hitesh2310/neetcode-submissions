func maxArea(heights []int) int {
l,r := 0,len(heights)-1
    res:=0
    for l<r {
        area :=  (r- l) * minOf(heights[l],heights[r])
        fmt.Println(l, r, area)
        if area > res {
            res = area
        } 
        if heights[l] < heights[r] {
            l++
        }else {
            r--
        }
    }
    return res
}

func maxOf(a,b int) int {
    if a > b {
        return a
    }
    return b 
}

func minOf(a,b int) int {
    if a < b {
        return a
    }
    return b 
}