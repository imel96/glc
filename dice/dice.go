package dice

import (
	"sort"
	"strconv"
	"strings"
)

func MinimumFaces(rolls []string) int {
	faces := 0

	for _, n := range max(transpose(createMatrix(rolls))) {
		faces += n
	}
	return faces
}

func createMatrix(rolls []string) [][]int {
	var matrix [][]int

	for _, roll := range rolls {
		var tmp []int

		for _, pip := range strings.Split(roll, "") {
			face, _ := strconv.Atoi(pip)
			tmp = append(tmp, face)
		}
		sort.Sort(sort.IntSlice(tmp))
		matrix = append(matrix, tmp)
	}
	return matrix
}

func transpose(matrix [][]int) [][]int {
	var tmatrix [][]int

	for i := 0; i < len(matrix[0]); i++ {
		var tmp []int

		for _, row := range matrix {
			tmp = append(tmp, row[i])
		}
		tmatrix = append(tmatrix, tmp)
	}
	return tmatrix
}

func max(matrix [][]int) []int {
	var mrow []int

	for _, row := range matrix {
		sort.Sort(sort.IntSlice(row))
		mrow = append(mrow, row[len(row) - 1])
	}
	return mrow
}
