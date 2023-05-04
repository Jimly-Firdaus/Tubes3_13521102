package FeatureStringmatching

func BuildLast(pattern string) [128]int {
	var last [128]int
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}

func BmMatch(text, pattern string) int {
	last := BuildLast(pattern)
	n, m := len(text), len(pattern)
	i := m - 1
	if i > n-1 {
		return -1
	}
	j := m - 1
	for i <= n-1 {
		if pattern[j] == text[i] {
			if j == 0 { // Match found
				if n == m {
					return -2 // EXACT MATCH
				} else {
					return i
				}
			} else {
				i--
				j--
			}
		} else {
			lo := last[text[i]]
			i += m - Min(j, 1+lo)
			j = m - 1
		}
	}
	return -1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
