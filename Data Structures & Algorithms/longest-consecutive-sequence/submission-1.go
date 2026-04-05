func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	
	history := map[int]int{}
	cumulativeMax := 0

	for _,num := range nums {
		val,ok := history[num-1]
		history[num] = 1

		if ok {
			history[num] += val
		}

		if history[num] > cumulativeMax {
			cumulativeMax = history[num]
		} 
	}

	return cumulativeMax 
	

}
