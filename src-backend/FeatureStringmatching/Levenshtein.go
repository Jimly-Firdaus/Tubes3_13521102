package FeatureStringmatching

func minList(numbers ...int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func levenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}
			matrix[i][j] = minList(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost,
			)
		}
	}
	return matrix[m][n]
}


// tester
// func main() {
// 	s := "aba"
// 	t := "bab"
// 	distance := levenshteinDistance(s, t)
// 	percentage := float64(1) - (float64(distance) / float64(len(s)))
// 	fmt.Printf("The Levenshtein distance between %q and %q is %f with %d mistakes\n", s, t, percentage, distance)
// }
