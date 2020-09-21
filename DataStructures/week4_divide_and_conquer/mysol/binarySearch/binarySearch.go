package main

import (
	"fmt"
)

type queryPos struct {
	pos, val, search int
}

//so this is the whole merge sort algorithm which can sort the elements in nlogn time period
func mergeSort(slice []queryPos, how string) []queryPos {
	if len(slice) == 1 {
		return slice
	}
	if how == "val" {
		slice = merge(mergeSort(slice[:len(slice)/2], "val"), mergeSort(slice[len(slice)/2:], "val"))
	} else {
		slice = mergeDifferent(mergeSort(slice[:len(slice)/2], "pos"), mergeSort(slice[len(slice)/2:], "pos"))
	}
	return slice
}

//merging function to merge two slices in a sorted order in linear time
func merge(slice1, slice2 []queryPos) []queryPos {
	slice := make([]queryPos, 0, len(slice1)+len(slice2))

	var counter1, counter2 int

	for counter1 != len(slice1) && counter2 != len(slice2) {
		if slice1[counter1].val < slice2[counter2].val {
			slice = append(slice, slice1[counter1])
			counter1++
		} else {
			slice = append(slice, slice2[counter2])
			counter2++
		}
	}

	if counter1 != len(slice1) {
		for counter1 != len(slice1) {
			slice = append(slice, slice1[counter1])
			counter1++
		}
	} else {
		for counter2 != len(slice2) {
			slice = append(slice, slice2[counter2])
			counter2++
		}
	}

	return slice
}

//merging function to merge two slices in a sorted order in linear time
func mergeDifferent(slice1, slice2 []queryPos) []queryPos {
	slice := make([]queryPos, 0, len(slice1)+len(slice2))

	var counter1, counter2 int

	for counter1 != len(slice1) && counter2 != len(slice2) {
		if slice1[counter1].pos < slice2[counter2].pos {
			slice = append(slice, slice1[counter1])
			counter1++
		} else {
			slice = append(slice, slice2[counter2])
			counter2++
		}
	}

	if counter1 != len(slice1) {
		for counter1 != len(slice1) {
			slice = append(slice, slice1[counter1])
			counter1++
		}
	} else {
		for counter2 != len(slice2) {
			slice = append(slice, slice2[counter2])
			counter2++
		}
	}

	return slice
}

//provided for taking input from user in the form of text than appending them to a slice and returning the slice
func inputSlice() (int, []int) {

	//initializing the input slice with length capacity
	var length int
	fmt.Scanf("%d", &length)
	slice := make([]int, 0, length)

	//taking inputs
	var val int
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &val)
		slice = append(slice, val)
	}
	return length, slice
}

func binarySearch(sortedSlice []int, toSearch []queryPos) {
	//declaring the start and end to go through the list
	var (
		start int = 0
		end   int = len(sortedSlice) - 1
		mid   int = -1
	)

	var saveMid int
	//iteraive version of binary search
	for i := 0; i < len(toSearch); i++ {
		saveMid = start
		for start <= end {
			mid = start + (end-start)/2
			if sortedSlice[mid] == toSearch[i].val {
				toSearch[i].search = mid
				break
			} else if sortedSlice[mid] < toSearch[i].val {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		end = len(sortedSlice) - 1
		if toSearch[i].search == -1 {
			start = saveMid
		} else {
			start = mid
		}
	}
}

func main() {
	//taking the sorted Slice and the queries as input
	_, sortedSlice := inputSlice()
	query, queries := inputSlice()

	// fmt.Println(sortedSlice)
	// fmt.Println(query, queries)

	posSliceQuery := make([]queryPos, 0, query)
	//sorting the input query slice to perform binary search more effeciently
	for i, v := range queries {
		posSliceQuery = append(posSliceQuery, queryPos{i, v, -1})
	}
	posSliceQuery = mergeSort(posSliceQuery, "val")

	// fmt.Println(posSliceQuery)

	binarySearch(sortedSlice, posSliceQuery)

	// fmt.Println(posSliceQuery)
	posSliceQuery = mergeSort(posSliceQuery, "pos")

	// fmt.Println(posSliceQuery)

	for i := 0; i < query; i++ {
		fmt.Print(posSliceQuery[i].search, " ")
	}
	fmt.Println()
}
