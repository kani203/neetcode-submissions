func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    for _, word := range strs {
        var cnt [26]int
        for _, c := range word {
            cnt[c - 'a']++
        }
        groups[cnt] = append(groups[cnt], word)
    }

    result := make([][]string, 0, len(groups))

    for _, words := range groups {
        result = append(result, words)
    }
    return result
}
