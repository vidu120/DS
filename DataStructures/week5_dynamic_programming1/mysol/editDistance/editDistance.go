package main

import (
	"fmt"
	"math"
)

func editDistanceTrack(firstString, secondString string, trace *[101][101]int) {
	//for tracing purposes

	//first fill up the known rows and columns
	//for first string
	for i := 0; i <= len(secondString); i++ {
		trace[i][0] = i
	}

	//for second string
	for i := 0; i <= len(firstString); i++ {
		trace[0][i] = i
	}
	//for comparing basis
	var min int
	//for finding the edit distance , a simple nested loop
	for i := 1; i <= len(firstString); i++ {
		for j := 1; j <= len(secondString); j++ {
			min = math.MaxInt64
			if firstString[i-1] == secondString[j-1] {
				min = int(math.Min(float64(trace[j-1][i-1]), float64(min)))
			} else {
				min = int(math.Min(float64(trace[j-1][i-1]+1), float64(min)))
			}
			min = int(math.Min(float64(trace[j][i-1]+1), float64(min)))
			min = int(math.Min(float64(trace[j-1][i]+1), float64(min)))
			trace[j][i] = min
		}
	}
}

func main() {
	//two strings for comparing
	var firstString, secondString string

	//taking inputs from string
	fmt.Scan(&firstString)
	fmt.Scan(&secondString)

	//keeping the trace in track
	var trace [101][101]int = [101][101]int{}

	//performing the trace operation
	editDistanceTrack(firstString, secondString, &trace)

	fmt.Println(trace[len(secondString)][len(firstString)])

}
