import "slices"
func threeSum(nums []int) [][]int {
	// dict: x * -1 : location
	keyMap := make(map[int][]int)
	for idx, num := range nums {
		keyMap[num] = append(keyMap[num], idx)
	}
	// add two sum then find key in dict
	i := 0
	res := make([][]int, 0)
	// loop: i: 0 -> -2, j: 1 -> -1
	for i < len(nums) - 1 {
		j := i + 1
		for j < len(nums) {
			tmpRes := getValidIdx(keyMap, nums, i, j)
			if len(tmpRes) != 0 && isNotDuplicate(tmpRes, res) {
				res = append(res, tmpRes)
			}
			j++
		}
		i++
	}
	return res
}

func getValidIdx(keyMap map[int][]int, nums []int, i int, j int) []int {
	res := make([]int, 0)
	if keyMap[-nums[i]-nums[j]] == nil {
		return res
	}
	for _, idx := range keyMap[-nums[i]-nums[j]] {
		if idx == i || idx == j {
			continue
		} else {
			return []int{nums[i], nums[j], nums[idx]}
		}
	}
	return res
}

func isNotDuplicate(subset []int, res [][]int) bool {
	for _, nums := range res {
		slices.Sort(nums)
		slices.Sort(subset)
		if slices.Equal(nums, subset) {
			return false
		}
	}
	return true
}