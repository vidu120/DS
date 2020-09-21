package main

import "fmt"

func inputSort(slice []int, quan int) []int {

	var input int
	fmt.Scanf("%d", &input)
	slice = append(slice, input)

	for i := 0; i < quan-1; i++ {
		fmt.Scanf("%d", &input)
		for j := 0; j < len(slice); j++ {
			if slice[j] <= input {
				slice = append(slice, 0)
				copy(slice[j+1:len(slice)], slice[j:len(slice)-1])
				slice[j] = input
				break
			} else if j == len(slice)-1 {
				slice = append(slice, input)
				break
			}
		}
	}

	return slice
}

func getMaxRevenue(adsRevenue, slotClicks []int) int {
	var revenue int

	for i := 0; i < len(adsRevenue); i++ {
		revenue += adsRevenue[i] * slotClicks[i]
	}

	return revenue
}

func main() {

	//number of ads available
	var numberOfAds int
	fmt.Scan(&numberOfAds)

	//slices initialize
	adsRevenue := make([]int, 0, numberOfAds)
	slotClicks := make([]int, 0, numberOfAds)

	//two slices for ads revenue and slot clicks
	adsRevenue = inputSort(adsRevenue, numberOfAds)
	slotClicks = inputSort(slotClicks, numberOfAds)

	fmt.Println(getMaxRevenue(adsRevenue, slotClicks))
}
