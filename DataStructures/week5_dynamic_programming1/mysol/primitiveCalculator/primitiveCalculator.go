//Based on dynamic programming

package main

import (
	"fmt"
	"math"
)

func minOperations(number int) []int {

	//declaring the backtracking slice
	backTrack := make([]int, 0, number-1)

	//For 1 number
	backTrack = append(backTrack, 0)

	//Main code
	var min int

	//iterative version
	for i := 1; i < number; i++ {
		//resetting the value
		min = math.MaxInt64

		//conditional for the soln
		if (i+1)%2 == 0 {
			min = int(math.Min(float64(backTrack[(i+1)/2-1]+1), float64(min)))
		}
		if (i+1)%3 == 0 {
			min = int(math.Min(float64(backTrack[(i+1)/3-1]+1), float64(min)))
		}
		min = int(math.Min(float64(backTrack[i-1]+1), float64(min)))

		//appending to the slice
		backTrack = append(backTrack, min)
	}

	return backTrack[:number]
}

func routeTrack(backTrack []int) []int {
	revertedTrack := make([]int, 0, len(backTrack)+1)

	//back counter for backtracking the optimal path
	var backCounter int = len(backTrack) - 1

	//reverted track for saving the backtracking
	revertedTrack = append(revertedTrack, len(backTrack))

	//loop for finding the optimal path
	for backCounter > 0 {
		if backTrack[backCounter] == backTrack[backCounter-1]+1 {
			revertedTrack = append(revertedTrack, backCounter)
			backCounter--
		} else if (backCounter+1)%2 == 0 && backTrack[backCounter] == backTrack[(backCounter+1)/2-1]+1 {
			revertedTrack = append(revertedTrack, (backCounter+1)/2)
			backCounter = (backCounter+1)/2 - 1
		} else {
			revertedTrack = append(revertedTrack, (backCounter+1)/3)
			backCounter = (backCounter+1)/3 - 1
		}
	}
	return revertedTrack
}

func main() {
	//take the number on which to perform the number of operations
	var number int
	fmt.Scan(&number)

	//Driver code...

	//Here we take the slice for backtracking purposes
	backtrack := minOperations(number)

	revertedTrack := routeTrack(backtrack)

	//solution for this problem
	fmt.Println(backtrack[len(backtrack)-1])
	for i := len(revertedTrack) - 1; i >= 0; i-- {
		fmt.Print(revertedTrack[i], " ")
	}
	fmt.Println()

}
