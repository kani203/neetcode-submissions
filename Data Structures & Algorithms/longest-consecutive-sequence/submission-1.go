func longestConsecutive(nums []int) int {
    cntMap := make(map[int]bool)
    for _, i := range nums {
        cntMap[i] = true
    }
    
    res := 0
    for i := range cntMap {
        if !cntMap[i-1] {
            curr := i
            tmpLen := 1
            for cntMap[curr+1] {
                curr++
                tmpLen++
            }
            if tmpLen > res {
                res = tmpLen
            }
        }
    }
    return res
}