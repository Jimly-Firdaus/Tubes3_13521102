package FeatureStringmatching

func buildLast(pattern string) [128]int {
	var last [128]int
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}

func bmMatch(text, pattern string) int {
	last := buildLast(pattern)
	n, m := len(text), len(pattern)
	i := m - 1
	if i > n-1 {
		return -1
	}
	j := m - 1
	for i <= n-1 {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[text[i]]
			i += m - min(j, 1+lo)
			j = m - 1
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// tester
// func main() {
// 	var text, pattern string
// 	text = "Apa Ibukota Indonesia"
// 	pattern = "Ibukota Indonesia"
// 	fmt.Println("Text:", text)
// 	fmt.Println("Pattern:", pattern)
// 	posn := bmMatch(text, pattern)
// 	if posn == -1 {
// 		fmt.Println("Pattern not found")
// 	} else {
// 		fmt.Println("Pattern starts at posn", posn)
// 	}
// }
