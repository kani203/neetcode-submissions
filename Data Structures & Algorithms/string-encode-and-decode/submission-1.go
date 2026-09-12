type Solution struct{}

const spliter byte = '#'

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res strings.Builder
	for _, word := range strs {
		res.WriteString(strconv.Itoa(len(word)))
		res.WriteByte(spliter)
		res.WriteString(word)
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i + 1
		for encoded[j] != spliter {
			j++
		}
		size, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, string(encoded[i:i+size]))
		i += size
	}
	return res
}
