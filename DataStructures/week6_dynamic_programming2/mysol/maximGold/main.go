package main

import (
	"fmt"
)

func inputSlice(slice []int) []int {
	var val int
	for i := 0; i < cap(slice); i++ {
		fmt.Scan(&val)
		slice = append(slice, val)
	}
	return slice
}

func maxGoldWeight(knapsackWeight int, goldBars []int, track *[301][10001]int) int {

	//main logic of this function , kinda complicated , but just another one of the use cases of dynamic programming
	var val int
	for i := 1; i <= len(goldBars); i++ {
		for j := 1; j <= knapsackWeight; j++ {
			track[i][j] = track[i-1][j]
			if goldBars[i-1] <= j {
				val = track[i-1][j-goldBars[i-1]] + goldBars[i-1]
				if val > track[i][j] {
					track[i][j] = val
				}
			}
		}
	}

	return track[len(goldBars)][knapsackWeight]
}

func main() {

	//first taking the input from the user
	var knapsackWeight, noOfBars int
	fmt.Scan(&knapsackWeight, &noOfBars)

	//golbars slice
	goldBars := make([]int, 0, noOfBars)
	goldBars = inputSlice(goldBars)

	var track [301][10001]int = [301][10001]int{}

	fmt.Println(maxGoldWeight(knapsackWeight, goldBars, &track))

}
