func twoSum(numbers []int, target int) []int {
    i := 0
    j := len(numbers) - 1
    for i < j && numbers[i] + numbers[j] != target {
        for numbers[i] + numbers[j] > target && i < j {
            j--
        }
        for numbers[i] + numbers[j] < target && i < j {
            i++
        }
    }
    return []int{i+1,j+1}
}
