//we have to use the concept of dynamic programming again

//Here we are using the same concept only the difference is of backtracking

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

func longestCommonSubSequence(first, second []int, trace *[101][101]int) {

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
}

func backTrack(trace *[101][101]int, firstSequence, secondSequence []int) []int {

	//make the common slice with the maximum capacity i.e 100
	common := make([]int, trace[len(secondSequence)][len(firstSequence)], 100)

	//for adding elements to the list
	backCounter := trace[len(secondSequence)][len(firstSequence)] - 1

	var row, column int

	row = len(secondSequence)
	column = len(firstSequence)

	//adding...
	for backCounter >= 0 {

		if firstSequence[column-1] == secondSequence[row-1] {
			common[backCounter] = firstSequence[column-1]
			backCounter--
		}

		if trace[row][column] == trace[row-1][column] {
			row--
		} else if trace[row][column] == trace[row][column-1] {
			column--
		} else {
			row--
			column--
		}
	}

	return common
}

func findMax(max int, trace *[101][101]int, first, second, third []int) int {
	longestCommonSubSequence(first, second, trace)

	//finding the common sequence using the trace left behind
	commonFirstSecond := backTrack(trace, first, second)

	//tracing for the newly obtained common sequence and the third sequence
	longestCommonSubSequence(commonFirstSecond, third, trace)

	if max < trace[len(third)][len(commonFirstSecond)] {
		max = trace[len(third)][len(commonFirstSecond)]
	}

	return max

}

func main() {

	//taking inputs of the sequences
	firstSequence := inputAsSlice()
	secondSequence := inputAsSlice()
	thirdSequence := inputAsSlice()
	// fmt.Println(firstSequence, secondSequence)

	var max int = 0

	var trace [101][101]int = [101][101]int{}

	//1
	max = findMax(max, &trace, firstSequence, secondSequence, thirdSequence)
	//2
	max = findMax(max, &trace, firstSequence, thirdSequence, secondSequence)
	//3
	max = findMax(max, &trace, secondSequence, thirdSequence, firstSequence)

	//longest common subsequence of the three sequences
	fmt.Println(max)
}
