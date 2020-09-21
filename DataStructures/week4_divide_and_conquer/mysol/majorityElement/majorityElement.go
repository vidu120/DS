package main

import "fmt"

//so this is the whole merge sort algorithm which can sort the elements in nlogn time period
func mergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}
	slice = merge(mergeSort(slice[:len(slice)/2]), mergeSort(slice[len(slice)/2:]))
	return slice
}

//merging function to merge two slices in a sorted order in linear time
func merge(slice1, slice2 []int) []int {
	slice := make([]int, 0, len(slice1)+len(slice2))

	var counter1, counter2 int

	for counter1 != len(slice1) && counter2 != len(slice2) {
		if slice1[counter1] < slice2[counter2] {
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

func majorityElement(slice []int) bool {

	//variables for keeping track of the frequency of an element
	var keepTrack, temp int
	keepTrack = slice[0]

	//loop for finding if such major element exists
	for _, val := range slice {
		if val == keepTrack {
			temp++
		} else {
			if temp > len(slice)/2 {
				return true
			}
			temp = 1
			keepTrack = val
		}
	}
	if temp > len(slice)/2 {
		return true
	}
	return false
}

//provided for taking input from user in the slice
func inputSlice(length int) []int {

	//initializing the input slice with length capacity
	slice := make([]int, 0, length)

	//taking inputs
	var val int
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &val)
		slice = append(slice, val)
	}
	return slice
}

func main() {
	//scanning input
	var length int
	fmt.Scanf("%d", &length)

	//printing the result
	if majorityElement(mergeSort(inputSlice(length))) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
