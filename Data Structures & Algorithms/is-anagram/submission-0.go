import (
	"maps"
)

func isAnagram(s string, t string) bool {
	sCharCount := make(map[rune]int)
	tCharCount := make(map[rune]int)

	for _,char := range s {
		_,ok := sCharCount[char]
	
		
		if ok {
			sCharCount[char] += 1
		} else {
			sCharCount[char] = 0
		}
		
	}

	for _,char := range t {
		_,ok  := tCharCount[char]
		
		if ok {
			tCharCount[char] += 1
		} else {
			tCharCount[char] = 0
		}
	}

	return maps.Equal(sCharCount,tCharCount)
 }