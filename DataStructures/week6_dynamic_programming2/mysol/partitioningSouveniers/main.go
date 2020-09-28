package main

import "fmt"

func inputSlice(slice []int) []int {

	var val int
	for i := 0; i < cap(slice); i++ {
		fmt.Scan(&val)
		slice = append(slice, val)
	}
	return slice
}

func ifPossible(slice []int) (bool, int) {
	var sum int
	for _, val := range slice {
		sum += val
	}
	if sum%3 == 0 && len(slice) >= 3 {
		return true, sum
	}
	return false, sum
}

func findSubsets(firstSubset, secondSubset, thirdSubset int, n int, slice []int) bool {

	if firstSubset == 0 && secondSubset == 0 && thirdSubset == 0 {
		return true
	}

	if n < 0 {
		return false
	}

	if findSubsets(firstSubset-slice[n], secondSubset, thirdSubset, n-1, slice) {
		return true
	} else if findSubsets(firstSubset, secondSubset-slice[n], thirdSubset, n-1, slice) {
		return true
	} else if findSubsets(firstSubset, secondSubset, thirdSubset-slice[n], n-1, slice) {
		return true
	}
	return false
}

func main() {

	var noOfSouvenier int
	fmt.Scan(&noOfSouvenier)

	//taking all the souvenier as input
	souvenier := make([]int, 0, noOfSouvenier)
	souvenier = inputSlice(souvenier)

	possibility, sum := ifPossible(souvenier)

	if possibility {
		if findSubsets(sum/3, sum/3, sum/3, len(souvenier)-1, souvenier) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else {
		fmt.Println(0)
	}

}
