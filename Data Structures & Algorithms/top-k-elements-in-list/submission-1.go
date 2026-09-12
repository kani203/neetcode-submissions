func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, w := range nums {
		counter[w]++
	}
	freq := make([][]int, len(nums)+1)
	for num, count := range counter {
		freq[count] = append(freq[count], num)
	}
	res := make([]int, 0, k)
	for j := len(nums); j > 0; j-- {
		for _, i := range freq[j] {
			res = append(res, i)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
