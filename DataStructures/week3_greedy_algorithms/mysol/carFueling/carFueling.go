package main

import "fmt"

func sortTheInput(input []int, numberOfStops, totalDist int) []int {

	var dist int
	input = append(input, totalDist)
	input = append(input, 0)

	for i := 0; i < numberOfStops; i++ {
		fmt.Scanf("%d", &dist)
		for j := 0; j < len(input); j++ {
			if input[j] <= dist || j == len(input)-1 {
				input = append(input, 0)
				copy(input[j+1:len(input)], input[j:len(input)-1])
				input[j] = dist
				break
			}
		}
	}
	fmt.Println(input)
	return input
}

func calculateMinimumRefills(stops []int, fuelDist int) int {
	var numberOfRefill, currentRefill, lastRefill int
	for currentRefill <= len(stops)-2 {
		lastRefill = currentRefill
		for currentRefill <= len(stops)-2 && stops[len(stops)-currentRefill-2]-stops[len(stops)-lastRefill-1] <= fuelDist {
			currentRefill = currentRefill + 1
		}
		if currentRefill == lastRefill {
			numberOfRefill = -1
			break
		} else if currentRefill <= len(stops)-2 {
			numberOfRefill++
		}
	}
	return numberOfRefill
}

func main() {
	//scanning of the input from user done
	var totalDist, fuelDist, numberOfStops int
	fmt.Scan(&totalDist)
	fmt.Scan(&fuelDist)
	fmt.Scan(&numberOfStops)

	//make a slice of stops
	stops := make([]int, 0, numberOfStops)

	//sort the input using insertion sort
	stops = sortTheInput(stops, numberOfStops, totalDist)

	//printing the answer to the screen
	fmt.Println(calculateMinimumRefills(stops, fuelDist))
}
