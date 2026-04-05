type kv struct {
	Key int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	kvsSlice := []kv{}
	result := []int{}

	for _,num := range nums {
		countMap[num] += 1
	}

	for k,v := range countMap {
		kvsSlice = append(kvsSlice,kv{k,v})
	}

	sort.Slice(kvsSlice, func(i,j int) bool {
		return kvsSlice[i].Value > kvsSlice[j].Value
	})

	for _,item := range kvsSlice{
		result = append(result,item.Key)
	}

	return result[:k]
}