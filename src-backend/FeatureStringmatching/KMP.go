package FeatureStringmatching

import "fmt"

func ComputeBorder(substr string) []int {
	table := make([]int, len(substr))
	table[0] = 0
	j := 0
	i := 1
	for i < len(substr) {
		if substr[j] == substr[i] {
			table[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = table[j-1]

		} else {
			table[i] = 0
			i++
		}
	}
	return table
}

func KMP(str string, substr string) int {
	len_str := len(str)
	len_sub := len(substr)
	table := ComputeBorder(substr)

	// Heuristic simple search
	if len_sub > len_str {
		return 0
	}
	if len_sub == len_str {
		if str == substr {
			return 1
		} else {
			return 0
		}
	}

	i, j := 0, 0
	for i < len_str {
		if substr[j] == str[i] {
			if j == len_sub-1 {
				return (i - len_sub + 1) // Match found
			}
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = table[j-1]
		}
		if len_sub == j {
			fmt.Printf("Disini kah?\n")
			return 1
		}
	}
	return -1
}
// Main tester
// func main() {
// 	s := "Bagaimana algoritma genetika bekerja?"
// 	t := "Bagaimana cara kerja algoritma genetika?"
//
// 	if KMP(t, s) == -1 {
// 		fmt.Printf("s is not a suffix\n")
// 	} else {
// 		fmt.Printf("s is a suffix.\n")
// 	}
// 	fmt.Printf("Starting index: %d", KMP(t, s))
// }
