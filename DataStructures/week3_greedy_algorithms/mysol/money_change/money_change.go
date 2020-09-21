package main

import "fmt"

func changeIntoDenominations(totalMoney *int) int {
	var totalCoins int = 0

	//denominations
	var denominations [3]int = [3]int{10, 5, 1}
	for i := 0; i < 3; i++ {
		totalCoins += *totalMoney / denominations[i]
		*totalMoney = *totalMoney % denominations[i]
	}
	return totalCoins
}

func main() {
	//declaring total money variable
	var totalMoney int

	//scanning for user input
	fmt.Scan(&totalMoney)

	totalCoinsUsed := changeIntoDenominations(&totalMoney)
	fmt.Println(totalCoinsUsed)
}
