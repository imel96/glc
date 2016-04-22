package dice

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func MinimumFaces(rolls []string) int {
	var max int

	for _, roll := range rolls {

		i := 0
		tmp := make([]int, len(roll));

		for _, pip := range strings.Split(roll, "") {
			face, _ := strconv.Atoi(pip)
			tmp[i] = face
			i++
		}
		sort.Sort(sort.IntSlice(tmp))
		fmt.Printf("pip %v\n", tmp)
	}
	return max
}
