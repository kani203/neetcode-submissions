func lengthOfLongestSubstring(s string) int {
	res := 0
	left := 0
	rs := []rune(s)
	cnt := make(map[rune]int)
	for right, ch := range rs {
		if idx, ok := cnt[ch]; ok && idx >= left {
			left = idx + 1
		}
		cnt[ch] = right
		if right-left+1 > res {
            res = right - left + 1
        }
	}
	return res
}
