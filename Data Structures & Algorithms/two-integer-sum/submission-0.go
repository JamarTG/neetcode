func twoSum(nums []int, target int) []int {
	history := map[int]int{}

	for i := 0; i < len(nums); i++ {
		x := target - nums[i]

		j,ok := history[x]
		
		if ok {
			return []int {j,i}
		}

		history[nums[i]] = i 
	}

	return []int{}	
}