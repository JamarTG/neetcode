
func isPalindrome(s string) bool {
	
	str := ""

	for _,char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			str += string(char)
		}
	}

	str = strings.ToLower(str)

	l := 0
	r := len(str)-1

	for l < r {
		fmt.Println(string(str[l]), string(str[r]))
		if str[l] != str[r] {
			return false
		}

		l++
		r--
	} 

	return true

}
