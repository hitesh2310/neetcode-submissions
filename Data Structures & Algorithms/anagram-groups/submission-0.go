func groupAnagrams(strs []string) [][]string {
    res := [][]string{}
    for _, str := range strs {
        fmt.Println("Current String::", str)
        if len(res) == 0 {
                tempArr := []string{}
                tempArr = append(tempArr,str)    
                res = append(res,tempArr)
            }else {
                matched := false
                 for index, eachArr := range res {
                    fmt.Println("Current Each Arr::",eachArr)
                   if len(str) != len(eachArr[0]) {
                        continue
                    }
                    charMapOfStr :=  makeCharMap(str)
                    charMapOfStrCmp := makeCharMap(eachArr[0])
                    if areTwoMapSame(charMapOfStr, charMapOfStrCmp) {
                        fmt.Println("Matched::", str,eachArr[0])
                        eachArr = append(eachArr,str)
                        fmt.Println(eachArr)
                        res[index] = eachArr
                        matched = true
                    }
                }
                if !matched {
                    tempArr := []string{}
                    tempArr = append(tempArr,str)
                    res = append(res,tempArr)
                }    
            }
    }

    return res
}




func makeCharMap(s string) map[rune]int{ 
        charMap := make(map[rune]int,26)
        for i:='a';i<='z';i++ {
            charMap[i] = 0
        }
        for _, char := range s {
          charMap[char] = charMap[char] + 1   
        }
    return charMap
}



func areTwoMapSame(a,b map[rune]int) bool {

    for key,val := range a {
        if b[key] != val {
            return false
        }
    }    

    return true    
}