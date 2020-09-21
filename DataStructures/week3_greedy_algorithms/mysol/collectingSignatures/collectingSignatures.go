package main

import "fmt"

//for our segments
type segment struct {
	start int
	end   int
}

type lastSavedSegment struct {
	lastSavedSegment segment
	usage            bool
}


//this is basically insertion sort , but not in a proper way , no need to copy the elements again and again
func inputSortedSegments(segments []segment, numberOfSegments int) []segment {
	//for inputting value from user
	var start, end int
	//inputting the first segment for reference
	fmt.Scanf("%d %d", &start, &end)
	segments = append(segments, segment{start: start, end: end})
	for i := 0; i < numberOfSegments-1; i++ {
		fmt.Scanf("%d %d", &start, &end)
		for j := 0; j < len(segments); j++ {
			if segments[j].start > start || (segments[j].start == start && segments[j].end >= end) {
				segments = append(segments, segment{})
				copy(segments[j+1:len(segments)], segments[j:len(segments)-1])
				segments[j] = segment{start: start, end: end}
				break
			} else if j == len(segments)-1 {
				segments = append(segments, segment{start: start, end: end})
				break
			}
		}
	}
	return segments
}

func compare(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
The below function was a failure cause it was not well formed , i just created it randomly without thinking thoroghly about the result
*/

// func minimumPoints(segments []segment) (int, []int) {
// 	var points int = 0
// 	var lastSaved lastSavedSegment = lastSavedSegment{lastSavedSegment: segment{}, usage: false}

// 	coordinates := make([]int, 0)

// 	//special case
// 	if len(segments) == 1 {
// 		return 1, []int{segments[0].start}
// 	}

// 	for i := 0; i < len(segments)-1; i++ {
// 		if segments[i+1].start > segments[i].end {
// 			points++
// 			coordinates = append(coordinates, segments[i].end)

// 			if lastSaved.usage {
// 				points++
// 				coordinates = append(coordinates, lastSaved.lastSavedSegment.end)
// 				lastSaved.usage = false
// 			}

// 			fmt.Println(lastSaved)
// 			// fmt.Println("Points : ", points, "temp : ", temp)
// 		} else {
// 			if lastSaved.usage {
// 				if lastSaved.lastSavedSegment.end < segments[i+1].start {
// 					points++
// 					coordinates = append(coordinates, lastSaved.lastSavedSegment.end)
// 					// fmt.Println("Points : ", points, "temp : ", temp)
// 					lastSaved.usage = false
// 					fmt.Println(lastSaved)
// 				} else {
// 					// fmt.Println("Points : ", points, "temp : ", temp)
// 					lastSaved.lastSavedSegment = segment{start: segments[i+1].start, end: compare(segments[i+1].end, segments[i].end)}
// 					if i == len(segments)-2 {
// 						coordinates = append(coordinates, lastSaved.lastSavedSegment.end)
// 						points++
// 					}
// 					fmt.Println(lastSaved)
// 				}
// 			} else {
// 				lastSaved.lastSavedSegment = segment{start: segments[i+1].start, end: compare(segments[i+1].end, segments[i].end)}
// 				lastSaved.usage = true
// 				if i == len(segments)-2 {
// 					coordinates = append(coordinates, lastSaved.lastSavedSegment.end)
// 					points++
// 				}
// 				fmt.Println(lastSaved)
// 			}
// 		}
// 	}
// 	return points, coordinates
// }

//this is a well formed function after carefully planning first in the copy
func minimumPoints(segments []segment) (int, []int) {
	//initializing the points variable and the slice of coordinates
	var points int
	pointsPos := make([]int, 0)

	//to keep track of the previous entry
	var lastSaved lastSavedSegment = lastSavedSegment{lastSavedSegment: segment{}, usage: false}

	for i := 0; i < len(segments); i++ {
		//the following if statement for the two cases which is if there was any last saved segment
		if lastSaved.usage {
			//the following if statement is for when we run out of segments and we are at the last one
			if i <= len(segments)-2 {
				//the following if statement is for when the next segment is in the current segment
				if segments[i+1].start <= segments[i].end {
					//the following if statement is for when the last saved coincides with the with the next saved
					if segments[i+1].start <= lastSaved.lastSavedSegment.end {
						lastSaved.lastSavedSegment = segment{start: segments[i+1].start, end: compare(lastSaved.lastSavedSegment.end, segments[i+1].end)}
					} else {
						points++
						pointsPos = append(pointsPos, lastSaved.lastSavedSegment.end)
						lastSaved.usage = false
					}
				} else {
					points++
					pointsPos = append(pointsPos, lastSaved.lastSavedSegment.end)
					lastSaved.usage = false
				}

			} else {
				points++
				pointsPos = append(pointsPos, lastSaved.lastSavedSegment.end)
				lastSaved.usage = false
			}
		} else {
			//for end of segments
			if i <= len(segments)-2 {
				//if the next segment lies in the first segment
				if segments[i+1].start <= segments[i].end {
					lastSaved.usage = true
					lastSaved.lastSavedSegment = segment{start: segments[i+1].start, end: compare(segments[i].end, segments[i+1].end)}
				} else {
					points++
					pointsPos = append(pointsPos, segments[i].end)
				}
			} else {
				points++
				pointsPos = append(pointsPos, segments[i].end)
			}
		}
	}

	//returning the number of points and their respective positions
	return points, pointsPos

}

func main() {

	//scanning the number of segments
	var numberOfSegments int
	fmt.Scan(&numberOfSegments)

	//declaring an empty segment slice
	segments := make([]segment, 0, numberOfSegments)

	//take input of segements and starting them according to the starting point
	segments = inputSortedSegments(segments, numberOfSegments)
	// fmt.Println(segments)

	//printing the minimum number of points to cover all segments and their coordinates
	numberOfPoints, coordinates := minimumPoints(segments)

	fmt.Println(numberOfPoints)
	for _, value := range coordinates {
		fmt.Print(value, " ")
	}
	fmt.Println()
}
