func maxArea(heights []int) int {

	l := 0
	r := len(heights)-1
	maxWaterAmount := 0

	for l < r {
		shorterBarSize := min(heights[l],heights[r])
		waterAmount := (r-l) *  shorterBarSize
		maxWaterAmount = max(maxWaterAmount,waterAmount)

		if shorterBarSize == heights[l] {
			l++
		} else {
			r--
		}
	}

	return maxWaterAmount
}	
