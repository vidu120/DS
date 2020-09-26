//we have to use the concept of dynamic programming again

package main

import (
	"fmt"
	"math"
)

func inputAsSlice() []int {
	var length int
	fmt.Scan(&length)

	slic := make([]int, 0, length)

	var val int
	for length > 0 {
		fmt.Scanf("%d", &val)
		slic = append(slic, val)
		length--
	}
	return slic
}

func longestCommonSubSequence(first, second []int) int {
	var trace [101][101]int = [101][101]int{}

	//for comparing basis
	var max int
	//for finding the maximum comman subsequence of these two squences
	for i := 1; i <= len(first); i++ {
		for j := 1; j <= len(second); j++ {
			max = math.MinInt64

			//main conditional for our logic , for finding the longest sequence
			if first[i-1] == second[j-1] {
				max = int(math.Max(float64(trace[j-1][i-1]+1), float64(max)))
			} else {
				max = int(math.Min(float64(trace[j-1][i-1]), float64(max)))
			}
			max = int(math.Max(float64(trace[j][i-1]), float64(max)))
			max = int(math.Max(float64(trace[j-1][i]), float64(max)))

			trace[j][i] = max
		}
	}

	return trace[len(second)][len(first)]
}

func main() {

	//taking inputs of the sequences
	firstSequence := inputAsSlice()
	secondSequence := inputAsSlice()
	// fmt.Println(firstSequence, secondSequence)

	fmt.Println(longestCommonSubSequence(firstSequence, secondSequence))
}
