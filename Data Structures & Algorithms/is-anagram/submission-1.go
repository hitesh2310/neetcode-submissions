func isAnagram(s string, t string) bool {

    if len(s)!=len(t) {
        return false
    }

    mapS := make(map[byte]int)
    mapT := make(map[byte]int)

    for i:=0;i<len(s);i++{
        mapS[s[i]] = mapS[s[i]] + 1 
        mapT[t[i]] = mapT[t[i]] + 1 
    }

    for key, val := range mapS {
        if val!=mapT[key] {
            return false
        }
    }
return true
}
