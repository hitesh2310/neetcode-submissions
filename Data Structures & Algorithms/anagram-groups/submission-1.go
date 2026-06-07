func groupAnagrams(strs []string) [][]string {
    res:= [][]string{}

  for _,str := range strs {
    if len(res) == 0 {
        temp := []string{}
        temp = append(temp,str)
        res = append(res,temp)
        continue
    }
    matchFound := false
    for  index , anagramArr := range res {
        strMap := makeCharMap(str)
        anagramMap := makeCharMap(anagramArr[0])
        if areSame(strMap,anagramMap) {
            anagramArr = append(anagramArr,str)
            res[index] = anagramArr
            matchFound = true
            break
        }
    }  
    
    if !matchFound {
        temp := []string{}
        temp = append(temp,str)
        res = append(res,temp)
    }

  }   
   return res
}



func makeCharMap(s string) map[rune]int {
    charMap := make(map[rune]int)
    for i:='a' ; i<='z';i++ {
        charMap[i] = 0
    }
    for _, char := range s {
        charMap[char]++
    }
return charMap
}


func areSame(map1, map2 map[rune]int) bool  {
    if len(map1)!=len(map2) {
        return false
    }

    for key,val := range map1 {
        if val!=map2[key] {
            return false
        }
    }

return true
}