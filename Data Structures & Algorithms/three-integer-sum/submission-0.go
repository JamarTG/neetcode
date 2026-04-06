func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    comb := [][]int{}
    seen := make(map[string]bool)

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        innerTarget := 0 - nums[i]
        m := make(map[int]int)

        for j := i + 1; j < len(nums); j++ {
            needed := innerTarget - nums[j]
            k, ok := m[needed]

            if ok {
                trip := []int{nums[i], nums[k], nums[j]}
                sort.Ints(trip)
                key := fmt.Sprintf("%d,%d,%d", trip[0], trip[1], trip[2])
                if !seen[key] {
                    comb = append(comb, trip)
                    seen[key] = true
                }
            }

            m[nums[j]] = j
        }
    }

    return comb
}