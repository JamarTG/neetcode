import "slices"

func hasDuplicate(nums []int) bool {
	m := []int {}

	for _,num :=  range nums {
	
		ok := slices.Contains(m,num)

		if ok {
			return true
		}

		m = append(m,num)
		
	}
	
	return false
}