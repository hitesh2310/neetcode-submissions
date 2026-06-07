func lengthOfLongestSubstring(s string) int {

   if len(s) <  2 {
    return len(s)
   }
    myMap := make(map[string]int)
    l,r := 0,0
    maxLength := 0
    for r<len(s) {
        currentLetter := string(s[r])
        if val, exists := myMap[currentLetter]; exists {
            fmt.Println("l::", l , " r::",r ," val::",val)
            l = maxOf(l,myMap[currentLetter]+1)
            fmt.Println("new l::",l)
        }
            myMap[currentLetter] = r
            maxLength = maxOf(maxLength,r-l+1)
        

        r++
    }
    
    return maxLength
}


func maxOf(a,b int) int {
    if a > b {
        return a
    }
    return b
}
