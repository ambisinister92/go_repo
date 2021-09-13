package counter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Count(text string, n int) []string {

	if n < 0 {
		fmt.Println("invalid syntax")
		return nil
	}

	text = strings.ToLower(text)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	sa := strings.FieldsFunc(text, f)
	words := map[string]int{}
	for i := 0; i < len(sa); i++ {
		words[sa[i]]++
	}

	keys := make([]string, 0, len(words))
	for key := range words {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})

	result := make([]string, n, n)
	for y := 0; y < len(result) && y < len(keys); y++ {
		result[y] = keys[y] + " : " + strconv.Itoa(words[keys[y]])
	}

	return result

}
