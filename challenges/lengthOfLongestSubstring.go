package challenges

import "strings"

func lengthOfLongestSubstring(s string) int {
	result := ""
	m := make(map[string]int)
	for _, c := range s {
		if strings.Contains(result, string(c)) {
			m[result] = len(result)
			result = strings.TrimLeft(result, strings.SplitAfterN(result, string(c), 2)[0])
			result += string(c)
		} else {
			result += string(c)
			m[result] = len(result)
		}
	}
	for s, i := range m {
		if i > len(result) {
			result = s
		}
	}
	return len(result)
}
