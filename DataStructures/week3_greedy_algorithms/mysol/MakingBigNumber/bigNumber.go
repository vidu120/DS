package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareTwo(first, second string) bool {
	return first+second <= second+first
}

func inputNumbersAndSort(numberOfNums int) string {
	//for inputting value from user
	var num int
	sortedNums := make([]string, 0, numberOfNums)

	convertToString := func(a int) string {
		return strconv.Itoa(a)
	}

	//inputting the first segment for reference
	fmt.Scanf("%d", &num)
	sortedNums = append(sortedNums, convertToString(num))

	for i := 0; i < numberOfNums-1; i++ {
		fmt.Scanf("%d", &num)
		for j := 0; j < len(sortedNums); j++ {
			if compareTwo(sortedNums[j], convertToString(num)) {
				sortedNums = append(sortedNums, "")
				copy(sortedNums[j+1:len(sortedNums)], sortedNums[j:len(sortedNums)-1])
				sortedNums[j] = convertToString(num)
				break
			} else if j == len(sortedNums)-1 {
				sortedNums = append(sortedNums, convertToString(num))
				break
			}
		}
	}
	return strings.Join(sortedNums, "")
}

func main() {

	//initialize variable numberOfnums to compare
	var numberOfNums int
	fmt.Scan(&numberOfNums)

	//printing the answer
	fmt.Println(inputNumbersAndSort(numberOfNums))
}
