package relationships

import "fmt"

func Zero(n int) [][]bool {
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}
	return matrix
}

func Print(source []string, relationship [][]bool) {
	minSourceNameLen := 0
	for _, s := range source {
		if len(s) > minSourceNameLen {
			minSourceNameLen = len(s)
		}
	}

	fmt.Printf("%-*s | ", minSourceNameLen, "")
	for i := range source {
		fmt.Printf("%-*s", minSourceNameLen, source[i])
		if i < len(source)-1 {
			fmt.Printf("| ")
		}
	}
	for i := range source {
		fmt.Println()
		fmt.Println("--------------------")
		fmt.Printf(" %-*s| ", minSourceNameLen, source[i])
		for j := range source {
			if relationship[i][j] {
				fmt.Printf("%-*s", minSourceNameLen, "1")
			} else {
				fmt.Printf("%-*s", minSourceNameLen, "0")
			}
			if j < len(source)-1 {
				fmt.Printf("| ")
			}
		}
	}
	fmt.Println()
}
