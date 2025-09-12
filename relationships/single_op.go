package relationships

import "slices"

func Transpose(a [][]bool) [][]bool {
	if len(a) == 0 {
		return nil
	}

	return foreachcell(len(a), func(i int, j int) bool {
		return a[j][i]
	})
}

func Complement(a [][]bool) [][]bool {
	if len(a) == 0 {
		return nil
	}

	return foreachcell(len(a), func(i int, j int) bool {
		return !a[i][j]
	})
}

func DefinitionDomain(a [][]bool) []int {
	if len(a) == 0 {
		return nil
	}

	res := make([]int, 0, len(a))
	for i := range a {
		for j := range a[0] {
			if i != j && a[i][j] && !slices.Contains(res, i) {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func MeaningDomain(a [][]bool) []int {
	if len(a) == 0 {
		return nil
	}

	res := make([]int, 0, len(a))
	for i := range a {
		for j := range a[0] {
			if i != j && a[i][j] && !slices.Contains(res, j) {
				res = append(res, j)
				break
			}
		}
	}
	return res
}

func BottomIntersection(a [][]bool, x int) []int {
	if len(a) == 0 || x < 0 || x >= len(a) {
		return nil
	}
	res := make([]int, 0, len(a[x]))
	for y := range a[x] {
		if a[x][y] {
			res = append(res, y)
		}
	}
	return res
}

func TopIntersection(a [][]bool, x int) []int {
	if len(a) == 0 || x < 0 || x >= len(a) {
		return nil
	}
	res := make([]int, 0, len(a))
	for y := range a {
		if a[y][x] {
			res = append(res, y)
		}
	}
	return res
}
