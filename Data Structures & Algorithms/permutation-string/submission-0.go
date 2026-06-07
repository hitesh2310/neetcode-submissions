func checkInclusion(s1 string, s2 string) bool {
    if len(s2)< len(s1){
        return false
    }

    l:=0
    chrMapS1 := makeCharMap(s1)
    chrMapS2 := makeCharMap("")
    
    for r,chr := range s2 {
        chrMapS2[chr] =  chrMapS2[chr]+1
       if r-l+1 == len(s1){
        fmt.Println(chrMapS1,chrMapS2)
        if compareMap(chrMapS1,chrMapS2){
            return true
        }
        chrMapS2[rune(s2[l])] =  chrMapS2[rune(s2[l])] - 1
        l++
       }
    }

    return false
}


func makeCharMap(s string) map[rune]int {
    chrMap := make(map[rune]int)
    for i:='a';i<='z';i++ {
        chrMap[i] = 0
    }

    for _, chr := range s {
        chrMap[chr] =  chrMap[chr]+1 
    }
    
    return chrMap
}


func compareMap(cm1, cm2 map[rune]int) bool {

    // if len(cm1) != len(cm2) {
    //     return false
    // }

    for key,val := range cm1 {
        if val != cm2[key] {
            return false
        }
    }

    return true
}
