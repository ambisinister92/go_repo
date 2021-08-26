package counter

import (

  "strings"
  "unicode"
  "sort"
  "strconv"
  
)

func Count(text string, n int) []string{



  text=strings.ToLower(text)
  //f := strings.Fields(text)
  f := func(c rune) bool {
    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
  }

  sa:=strings.FieldsFunc(text, f)
  words:= map[string]int{}
  for i:=0; i<len(sa); i++{
    words[sa[i]]++
  }


  keys:=make([]string,0,len(words))
  for key := range words{
    keys=append(keys,key)
  }
  sort.Slice(keys, func(i,j int)bool{
    return words[keys[i]]>words[keys[j]]
  })


  result:=make([]string,n,n)
  for y:=0; y<len(result);y++{
    result[y]=keys[y] +" : "+ strconv.Itoa(words[keys[y]])
  }

  return result

}
