func productExceptSelf(nums []int) []int {
	size := len(nums)
	prePro := make([]int, size)
	prePro[0] = 1
	surPro := make([]int, size)
	surPro[size - 1] = 1
	// prePro[i] = product of nums[0...i-1]
	// surPro[i] = product of nums[i+1...n-1]
	// res[i] = prePro[i] * surPro[i]
	for i := 1; i < size; i++ {
		if i == 1 {
			prePro[i] = nums[i - 1]
		} else {
			prePro[i] = prePro[i - 1] * nums[i - 1]
		}
	}
	for i := size - 2; i >= 0; i-- {
		if i == size - 2 {
			surPro[i] = nums[i + 1]
		} else {
			surPro[i] = surPro[i + 1] * nums[i + 1]
		}
	}
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = prePro[i] * surPro[i]
	}
	return res
}
