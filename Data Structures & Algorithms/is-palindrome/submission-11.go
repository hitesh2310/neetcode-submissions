func isPalindrome(s string) bool {
    s = strings.ReplaceAll(s,"?","")
    s = strings.ReplaceAll(s," ","")
    s = strings.ReplaceAll(s,",","")
    s = strings.ReplaceAll(s,"'","")
    s = strings.ReplaceAll(s,"\\`","")
    s = strings.ReplaceAll(s,".","")
    s = strings.ReplaceAll(s,":","")
    s = strings.ToLower(s)
    l := 0
    r := len(s) - 1 
    for l<r {
      if strings.ToLower(string(s[l]))!=strings.ToLower(string(s[r])){
            return false
        }
        l++
        r--
    }
    return true
}
