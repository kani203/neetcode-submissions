func threeSum(nums []int) [][]int {
	res := map[[3]int]struct{}{}
	sort.Ints(nums)
	keymap := make(map[int][]int)
	for idx, num := range nums {
		keymap[num] = append(keymap[num], idx)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmpRes, found := getValidRes(keymap, i, j, nums)
			if found {
				res[tmpRes] = struct{}{}
			}
		}
	}
	var result [][]int
	for triplet := range res {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}
	return result
}

func getValidRes(keymap map[int][]int, i int, j int, nums []int) ([3]int, bool) {
	if keymap[-nums[i]-nums[j]] == nil {
		return [3]int{}, false
	}
	for _, idx := range keymap[-nums[i]-nums[j]] {
		if idx == i || idx == j {
			continue
		} else {
			res := [3]int{nums[i], nums[j], nums[idx]}
			sort.Ints(res[:])
			return res, true
		}
	}
	return [3]int{}, false
}