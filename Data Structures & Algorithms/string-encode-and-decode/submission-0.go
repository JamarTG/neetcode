type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	result := ""

	for _,str := range strs {
        strlen := strconv.Itoa(len(str))
		result += strlen + "#" + str  
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
    idx := 0
    result := []string{}

    for encoded != ""  {
        char := string(encoded[idx])

        if char == "#" {
            digits := encoded[:idx]
            numToRead, _ := strconv.Atoi(digits)
            word := encoded[idx+1:idx+1+numToRead] 
            result = append(result,word)
            encoded = encoded[idx+1+numToRead:]
            idx = 0
        } else {
            idx++
        }
    }

    return result
}