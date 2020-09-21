package main

import (
	"fmt"
)

type item struct {
	val    int
	weight int
	rate   float64
}

//sorting using the insertion sort operation
func inputSort(sortedItems []item, numberOfItems int) []item {
	var val, weight int
	for i := 0; i < numberOfItems; i++ {
		fmt.Scanf("%d %d", &val, &weight)
		for j := 0; j < len(sortedItems); j++ {
			if sortedItems[j].rate <= float64(val)/float64(weight) || j == len(sortedItems)-1 {
				sortedItems = append(sortedItems, item{})
				copy(sortedItems[j+1:len(sortedItems)], sortedItems[j:len(sortedItems)-1])
				sortedItems[j] = item{val: val, weight: weight, rate: float64(val) / float64(weight)}
				break
			}
		}
	}
	return sortedItems
}

//finding the max loot value
func findMaxLoot(items []item, weightToAchieve int) float64 {
	var maxLootValue float64 = 0
	var counter int = 0
	for weightToAchieve != 0 && counter < len(items) {
		if items[counter].weight >= weightToAchieve {
			maxLootValue += float64(weightToAchieve) * items[counter].rate
			weightToAchieve -= weightToAchieve
		} else {
			maxLootValue += float64(items[counter].val)
			weightToAchieve -= items[counter].weight
		}
		counter = counter + 1
	}
	return maxLootValue
}

func main() {

	//taking number of items and weight of backpack as input from user
	var numberOfItems, weightToAchieve int
	fmt.Scanf("%d %d", &numberOfItems, &weightToAchieve)

	//defining an array of items and some variables
	var val, weight int
	sortedItems := make([]item, 0, numberOfItems)

	//taking in the first item for reference
	fmt.Scan("%d %d", &val, &weight)
	sortedItems = append(sortedItems, item{val: val, weight: weight, rate: float64(val) / float64(weight)})

	//scanning items from user and sorting at the same time using function
	sortedItems = inputSort(sortedItems, numberOfItems)
	// fmt.Println(sortedItems)

	//finding the max loot
	fmt.Printf("%.4f\n", findMaxLoot(sortedItems[:len(sortedItems)], weightToAchieve))
}
