func isAnagram(s string, t string) bool {
    if len(s)!=len(t) {
        return false
    }
    Map1 := make(map[rune]int)
    Map2 := make(map[rune]int)
    for _ ,str := range s {
        Map1[str] = Map1[str]+1
    }
    for _ ,str := range t {
        Map2[str] = Map2[str]+1
    } 


    for key,val := range Map1 {

        if val!=Map2[key] {
            return false
        }
    }
    return true
}
