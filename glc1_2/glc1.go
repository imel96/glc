package main

import (
	"fmt"
	"sort"
)

func main() {
	altitude := []int{30, 20, 20, 10}
	fmt.Println("Cost, 世界, ", altitude, minimumCost(altitude))

	altitude = []int{5, 7, 3}
	fmt.Println("Cost, 世界, ", altitude, minimumCost(altitude))

	altitude = []int{6, 8, 5, 4, 7, 4, 2, 3, 1}
	fmt.Println("Cost, 世界, ", altitude, minimumCost(altitude))

	altitude = []int{749, 560, 921, 166, 757, 818, 228, 584, 366, 88}
	fmt.Println("Cost, 世界, ", altitude, minimumCost(altitude))

	altitude = []int{712, 745, 230, 200, 648, 440, 115, 913, 627, 621, 186, 222, 741, 954, 581, 193, 266, 320, 798, 745}
	fmt.Println("Cost, 世界, ", altitude, minimumCost(altitude))

	names := []string{"Toshi", "Jun", "Teru", "Chihiro"}
	fmt.Println("Names, 世界, ", names)

	votes := []string{"Jun", "Chihiro", "Toshi", "Toshi"}
	fmt.Println("Votes, 世界, ", votes, getTheBestEngineer(names, votes))

	votes = []string{"Teru", "Teru", "Jun", "Jun"}
	fmt.Println("Votes, 世界, ", votes, getTheBestEngineer(names, votes))

	votes = []string{"Toshi", "Toshi", "Jun", "Jun"}
	fmt.Println("Votes, 世界, ", votes, getTheBestEngineer(names, votes))

	names = []string{"Toshi", "Jun"}
	fmt.Println("Names, 世界, ", names)

	votes = []string{"Toshi", "Jun"}
	fmt.Println("Votes, 世界, ", votes, getTheBestEngineer(names, votes))

	names = []string{"PhpLove", "phplove", "phpLove", "Phplove"}
	fmt.Println("Names, 世界, ", names)

	votes = []string{"phpLove", "phpLove", "phpLove", "PhpLove"}
	fmt.Println("Votes, 世界, ", votes, getTheBestEngineer(names, votes))

	M := 2
	heights := []int{1, 2, 1, 4, 3}

	fmt.Println("Heights, 世界, ", M, heights, minFloors(M, heights))

	M = 3
	heights = []int{1, 3, 5, 2, 1}

	fmt.Println("Heights, 世界, ", M, heights, minFloors(M, heights))

	M = 1
	heights = []int{43, 19, 15}

	fmt.Println("Heights, 世界, ", M, heights, minFloors(M, heights))

	M = 3
	heights = []int{19, 23, 9, 12}

	fmt.Println("Heights, 世界, ", M, heights, minFloors(M, heights))

	M = 12

	heights = []int{25, 18, 38, 1, 42, 41, 14, 16, 19, 46, 42, 39, 38, 31, 43, 37, 26, 41, 33, 37, 45, 27, 19, 24, 33, 11, 22, 20, 36, 4, 4}

	fmt.Println("Heights, 世界, ", M, heights, minFloors(M, heights))
}

func minimumCost(altitude []int) int {
	last := 0
	cost := 0

	for _, value := range altitude {

		if value <= last || last == 0 {
			last = value

		} else {
			cost += value - last
		}
	}
	return cost
}

func getTheBestEngineer(names []string, votes []string) string {
	max := 0
	winner := ""
	voteCounts := make(map[string]int)

	for index, name := range votes {

		if names[index] != name {
			voteCounts[name]++
		}
	}
	for _, votes := range voteCounts {

		if max < votes {
			max = votes
		}
	}
	for name, votes := range voteCounts {

		if max == votes {

			if winner == "" {
				winner = name

			} else {
				winner = ""
				break
			}
		}
	}
	return winner
}

func minFloors(M int, heights []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	rets := minFloorsArray(M, heights)

	sort.Sort(sort.Reverse(sort.IntSlice(rets)))
	return rets[len(rets)-1]
}

func minFloorsArray(M int, heights []int) []int {
	n := 0

	if len(heights) <= 1 || M <= 1 {
		return []int{n}
	}
	floors := make(map[int]int)

	for _, name := range heights {
		floors[name]++
	}
	for _, count := range floors {

		if count == M {
			return []int{n}
		}
	}
	for m := 1; m < M; m++ {
		n += heights[0] - heights[m]
	}
	if M >= len(heights) {
		return []int{n}
	}
	return append([]int{n}, minFloorsArray(M, heights[1:])...)
}
