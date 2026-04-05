func groupAnagrams(strs []string) [][]string {
	
	hashMap := map[string][]string {} 
	var result [][]string;

	for _,str := range strs {
		strChars := []rune(str)
		
		sort.Slice(strChars, func (i,j int ) bool {
			return strChars[i] < strChars[j]
		})

		key := string(strChars)

		hashMap[key] = append(hashMap[key],str)
		
	}

	for _,val := range hashMap {
		result = append(result,val)
	}

	return result
}