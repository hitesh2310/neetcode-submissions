func lengthOfLongestSubstring(s string) int {
    res:=0 
    if len(s) < 2 {
        return len(s)
    }
    myMap := make(map[byte]int)
    l := 0
    for r:=0;r<len(s);r++ {
        if index, exists := myMap[s[r]];exists {
            res = maxOf(res,r-l)
            l = maxOf(index+1,l)
        }
            myMap[s[r]] = r
            res = maxOf(res,r-l+1)
    }
return res
}

func maxOf(a,b int) int {
    if a > b {
        return a 
    }
    return b
}