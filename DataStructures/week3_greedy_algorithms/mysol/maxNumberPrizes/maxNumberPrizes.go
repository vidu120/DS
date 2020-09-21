package main

import (
	"fmt"
	"math"
)

func maximizeWinning(totalCandies int) {
	//declaring and initialzing the variables and main body
	var limit int = int((math.Sqrt(float64(8*totalCandies+1)) - 1) / 2)
	var extraValue bool = (limit*limit+limit)/2 < totalCandies

	//printing the result
	fmt.Println(limit)
	if extraValue {
		for i := 0; i < limit-1; i++ {
			fmt.Print(i+1, " ")
		}
		fmt.Print(limit + (totalCandies - ((limit*limit + limit) / 2)))
	} else {
		for i := 0; i < limit; i++ {
			fmt.Print(i+1, " ")
		}
	}
	fmt.Println()
}

func main() {

	//total number of candies
	var totalCandies int
	fmt.Scan(&totalCandies)

	//maximum winning for every child and also printing the result at the same time
	maximizeWinning(totalCandies)
}
