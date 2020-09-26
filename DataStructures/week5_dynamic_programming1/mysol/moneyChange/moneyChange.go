package main

import (
	"fmt"
	"math"
)

func minDenominations(money int) int {
	//storing the min required coin denominations in a slice of length money
	dynamicStorage := make([]int, 0, money+1)

	dynamicStorage = append(dynamicStorage, 0)
	//driver code
	var min int = math.MaxInt64
	for i := 1; i <= money; i++ {
		min = math.MaxInt64
		for _, coin := range [3]int{1, 3, 4} {
			if i >= coin {
				min = int(math.Min(float64(dynamicStorage[i-coin]+1), float64(min)))
			}
		}
		dynamicStorage = append(dynamicStorage, min)
	}

	return dynamicStorage[len(dynamicStorage)-1]
}

func main() {

	//take total money as input and the coin denominations as 1 , 3  , 4
	var money int
	fmt.Scan(&money)

	//printing the answer to stdout stream
	fmt.Println(minDenominations(money))

}
