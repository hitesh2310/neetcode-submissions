func isPalindrome(s string) bool {


s = strings.ReplaceAll(s,"?","")
s = strings.ReplaceAll(s," ","")
s = strings.ReplaceAll(s,",","")
s = strings.ReplaceAll(s,"'","")
s = strings.ReplaceAll(s,"\\`","")
s = strings.ReplaceAll(s,".","")
s = strings.ReplaceAll(s,":","")


fmt.Println(s)
l,r:=0,len(s)-1
for l<r {
    fmt.Println(strings.ToLower(string(s[l])), strings.ToLower(string(s[r])) , strings.ToLower(string(s[l]))==string(s[r]))
    if  strings.ToLower(string(s[l]))!=strings.ToLower(string(s[r])){
        return false
    }
    l++
    r--
}
return true
}
