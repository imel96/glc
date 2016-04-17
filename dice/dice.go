package dice

import (
	//"fmt"
	"strconv"
	"strings"
)

func MinimumFaces(rolls []string) int {
	var max int

	for _, roll := range rolls {

		for _, pip := range strings.Split(roll, "") {
			//fmt.Printf("pip %v\n", pip)
			face, _ := strconv.Atoi(pip)
			if max < face {
				max = face
			}
		}
	}
	return max
}
