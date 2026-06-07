func characterReplacement(s string, k int) int {

chrMap := make(map[byte]int)
l,res,maxF := 0,0,0
for r:=0;r<len(s);r++ {
    chrMap[s[r]] = chrMap[s[r]] + 1 
    
    if chrMap[s[r]] > maxF {
        maxF = chrMap[s[r]]
    }

    for r-l+1 - maxF > k {
        chrMap[s[l]]--
        l++
    }

    if r - l + 1 > res {
        res = r - l + 1
    }
}

return res 
}
