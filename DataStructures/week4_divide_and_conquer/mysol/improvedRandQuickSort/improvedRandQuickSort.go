package main

import (
	"fmt"
	"math/rand"
)

//this is optimized quick sort for slices containing same element repeatedly
func partition(slice []int) (int, int) {
	//declaring some important variables
	// defer fmt.Println("Orig -", slice)
	var (
		x         int = slice[0]
		j         int
		keepTrack int = -1
	)
	for i := 1; i < len(slice); i++ {
		if slice[i] == x {
			j++
			slice[i], slice[j] = slice[j], slice[i]
			if keepTrack < 0 {
				keepTrack = j
			}
		} else if slice[i] < x {
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

func quickSort(slice []int) {
	//recursive format
	if len(slice) <= 1 {
		return
	}

	//using random number for better result
	randm := rand.Intn(len(slice))
	slice[0], slice[randm] = slice[randm], slice[0]

	//we get the indexes for the equal elements
	m1, m2 := partition(slice)
	quickSort(slice[:m1])
	quickSort(slice[m2+1:])
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

	//peforming the quick sort operation
	slice := inputSlice(length)
	quickSort(slice)
	for _, val := range slice {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
