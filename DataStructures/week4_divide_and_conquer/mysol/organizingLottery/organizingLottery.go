package main

import (
	"fmt"
	"math"
	"math/rand"
)

type segment struct {
	start, end int
}

type tracePoints struct {
	pos, val, payoff int
}

//this is basically insertion sort , but not in a proper way , no need to copy the elements again and again
func sortingSegments(segments, forwardSorted, backwardSorted []segment, numberOfSegments int) ([]segment, []segment) {

	if numberOfSegments == 1 {
		return segments, segments
	}

	//inputting the first segment for reference
	//for forward sorting
	forwardSorted = append(forwardSorted, segment{start: segments[0].start, end: segments[0].end})

	for i := 1; i < numberOfSegments; i++ {
		for j := 0; j < len(forwardSorted); j++ {
			if forwardSorted[j].start > segments[i].start || (forwardSorted[j].start == segments[i].start && forwardSorted[j].end >= segments[i].end) {
				forwardSorted = append(forwardSorted, segment{})
				copy(forwardSorted[j+1:len(forwardSorted)], forwardSorted[j:len(forwardSorted)-1])
				forwardSorted[j] = segment{start: segments[i].start, end: segments[i].end}
				break
			} else if j == len(forwardSorted)-1 {
				forwardSorted = append(forwardSorted, segment{start: segments[i].start, end: segments[i].end})
				break
			}
		}
	}

	backwardSorted = append(backwardSorted, segment{start: segments[0].start, end: segments[0].end})

	for i := 1; i < numberOfSegments; i++ {
		for j := 0; j < len(backwardSorted); j++ {
			if backwardSorted[j].end > segments[i].end || (backwardSorted[j].end == segments[i].end && backwardSorted[j].start >= segments[i].start) {
				backwardSorted = append(backwardSorted, segment{})
				copy(backwardSorted[j+1:len(backwardSorted)], backwardSorted[j:len(backwardSorted)-1])
				backwardSorted[j] = segment{start: segments[i].start, end: segments[i].end}
				break
			} else if j == len(backwardSorted)-1 {
				backwardSorted = append(backwardSorted, segment{start: segments[i].start, end: segments[i].end})
				break
			}
		}
	}
	return forwardSorted, backwardSorted
}

func inputSlice(slice []tracePoints, length int) []tracePoints {

	var val int
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &val)
		slice = append(slice, tracePoints{val: val, pos: i, payoff: 0})
	}

	return slice
}

func get(slice []tracePoints, index int, which string) int {
	if which == "pos" {
		return slice[index].pos
	}
	return slice[index].val
}

//this is optimized quick sort for slices containing same element repeatedly
func partition(slice []tracePoints, which string) (int, int) {
	//declaring some important variables
	// defer fmt.Println("Orig -", slice)
	var (
		x         int = get(slice, 0, which)
		j         int
		keepTrack int = -1
	)
	for i := 1; i < len(slice); i++ {
		if get(slice, i, which) == x {
			j++
			slice[i], slice[j] = slice[j], slice[i]
			if keepTrack < 0 {
				keepTrack = j
			}
		} else if get(slice, i, which) < x {
			j++
			slice[i], slice[j] = slice[j], slice[i]
			if keepTrack > 0 {
				slice[keepTrack], slice[j] = slice[j], slice[keepTrack]
				keepTrack++
			}
		}
	}
	if keepTrack > 0 {
		slice[0], slice[keepTrack-1] = slice[keepTrack-1], slice[0]
		// fmt.Println("partition -", slice, keepTrack-1, j)
		return keepTrack - 1, j
	}
	slice[0], slice[j] = slice[j], slice[0]
	// fmt.Println("partition -", slice, j, j)
	return j, j
}

//quickSort 1st type
func quickSort(slice []tracePoints, which string) {
	//recursive format
	if len(slice) <= 1 {
		return
	}

	//using random number for better result
	randm := rand.Intn(len(slice))
	slice[0], slice[randm] = slice[randm], slice[0]

	//we get the indexes for the equal elements
	m1, m2 := partition(slice, which)
	quickSort(slice[:m1], which)
	quickSort(slice[m2+1:], which)
}

func inputSegments(length int) []segment {
	input := make([]segment, 0, length)
	var start, end int
	for i := 0; i < length; i++ {
		fmt.Scanf("%d %d", &start, &end)
		input = append(input, segment{start: start, end: end})
	}
	return input
}

//logic of the question
func coveringSegments(forwardSorted, backwardSorted []segment, newPoints []tracePoints) {

	var pos1, pos2 int

	//initializers for forward and backward
	startF := 0
	endF := len(forwardSorted) - 1
	startB := 0
	endB := len(backwardSorted) - 1

	for i := 0; i < len(newPoints); i++ {
		//first get both the points
		pos1 = binarySearchSegmentsF(forwardSorted, startF, endF, newPoints[i].val)
		pos2 = binarySearchSegmentsB(backwardSorted, startB, endB, newPoints[i].val)

		// fmt.Println(pos1, pos2)
		if pos1 == -1 || pos2 == -1 {
			newPoints[i].payoff = 0
			continue
		}

		newPoints[i].payoff = int(math.Max(0, float64(pos1-pos2+1)))
		startF = pos1
		startB = pos2
	}
}

func binarySearchSegmentsF(segments []segment, start, end, point int) int {
	var pos int = -1
	var equalPos int = -1
	var mid int

	// fmt.Println(segments[start : end+1])

	for start <= end {
		mid = start + (end-start)/2
		// fmt.Println(segments[mid].start, point, mid)
		if segments[mid].start == point {
			equalPos = mid
			start = mid + 1
		} else if segments[mid].start < point {
			pos = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if equalPos != -1 {
		return equalPos
	}
	return pos
}

func binarySearchSegmentsB(segments []segment, start, end, point int) int {
	var pos int = -1
	var equalPos int = -1
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if segments[mid].end == point {
			equalPos = mid
			end = mid - 1
		} else if segments[mid].end > point {
			pos = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if equalPos != -1 {
		return equalPos
	}
	return pos
}

//driver function
func main() {
	//no of segments and points
	var noOfSegments, noOfPoints int
	fmt.Scanf("%d %d", &noOfSegments, &noOfPoints)

	//declaring the segments and the points
	forwardSorted := make([]segment, 0, noOfSegments)
	backwardSorted := make([]segment, 0, noOfSegments)

	//sorting the segments and the points by insertion sort and quick sort algorithms
	forwardSorted, backwardSorted = sortingSegments(inputSegments(noOfSegments), forwardSorted, backwardSorted, noOfSegments)

	newPoints := make([]tracePoints, 0, noOfPoints)
	// quickSort(points)
	newPoints = inputSlice(newPoints, noOfPoints)

	//sort the order on the basis of the value
	quickSort(newPoints, "val")

	//final result for all segments
	coveringSegments(forwardSorted, backwardSorted, newPoints)

	//get the same order again
	quickSort(newPoints, "pos")

	for _, val := range newPoints {
		fmt.Print(val.payoff, " ")
	}
	fmt.Println()
	// fmt.Println(forwardSorted)
	// fmt.Println(backwardSorted)
	// fmt.Println(points)

}
