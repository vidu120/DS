package main

import (
	"fmt"
	"math"
)

//Point ...This is a point struct
type Point struct {
	x, y int
}

type pointsGroup struct {
	pointWithSameX []Point
}

//trying out the class functionality
func (point1 *Point) calculateDistance(point2 Point) float64 {
	return math.Sqrt(math.Pow(float64(point1.x-point2.x), 2) + math.Pow(float64(point1.y-point2.y), 2))
}

func mergeSort(points []Point) {
	if len(points) == 1 {
		return
	}
	//sorting the two parts of the array , the left and right
	mergeSort(points[:len(points)/2])
	mergeSort(points[len(points)/2:])

	//meging the two sorted parts
	merge(points[:len(points)/2], points[len(points)/2:])

}

func merge(left, right []Point) {
	//declaring newPoints slice then we will append the elements one by one in this
	newPoints := make([]Point, 0, len(left)+len(right))
	var counter1, counter2 int

	for counter1 != len(left) && counter2 != len(right) {
		if left[counter1].x < right[counter2].x || (left[counter1].x == right[counter2].x && left[counter1].y <= right[counter2].y) {
			newPoints = append(newPoints, left[counter1])
			counter1++
		} else {
			newPoints = append(newPoints, right[counter2])
			counter2++
		}
	}

	//adding the rest of the elements in the newPoints slice
	for counter1 != len(left) {
		newPoints = append(newPoints, left[counter1])
		counter1++
	}
	for counter2 != len(right) {
		newPoints = append(newPoints, right[counter2])
		counter2++
	}

	//copying it the original slice
	copy(left, newPoints[:len(newPoints)/2])
	copy(right, newPoints[len(newPoints)/2:])
}

func mergeDifferent(left, right []Point) []Point {
	//declaring newPoints slice then we will append the elements one by one in this
	newPoints := make([]Point, 0, len(left)+len(right))
	var counter1, counter2 int

	for counter1 != len(left) && counter2 != len(right) {
		if left[counter1].y < right[counter2].y || (left[counter1].y == right[counter2].y && left[counter1].x <= right[counter2].x) {
			newPoints = append(newPoints, left[counter1])
			counter1++
		} else {
			newPoints = append(newPoints, right[counter2])
			counter2++
		}
	}

	//adding the rest of the elements in the newPoints slice
	for counter1 != len(left) {
		newPoints = append(newPoints, left[counter1])
		counter1++
	}
	for counter2 != len(right) {
		newPoints = append(newPoints, right[counter2])
		counter2++
	}
	return newPoints
}

func findMinDist(points []Point) float64 {
	//min answer
	var min float64
	if len(points) == 1 {
		return math.MaxFloat64
	}
	for i := 0; i < len(points)-1; i++ {
		min = math.Min(math.MaxFloat64, points[i].calculateDistance(points[i+1]))
	}
	return min
}

func inputPoints(points []Point, length int) []Point {
	var x, y int
	for length > 0 {
		fmt.Scanf("%d %d", &x, &y)
		points = append(points, Point{x, y})
		length--
	}
	return points
}

func breakThemIntoSections(points []Point) []pointsGroup {
	//make a slice of points group
	pointsG := make([]pointsGroup, 0, len(points))

	//append the first point from it
	pointsG = append(pointsG, pointsGroup{pointWithSameX: []Point{points[0]}})

	//doing a linear run through points to append them in different groups
	for i := 1; i < len(points); i++ {
		if points[i].x != points[i-1].x {
			pointsG = append(pointsG, pointsGroup{pointWithSameX: []Point{points[i]}})
		} else {
			pointsG[len(pointsG)-1].pointWithSameX = append(pointsG[len(pointsG)-1].pointWithSameX, points[i])
		}
	}

	return pointsG
}

func divideAndConquer(pointsG []pointsGroup) float64 {
	if len(pointsG) == 1 {
		return findMinDist(pointsG[0].pointWithSameX)
	}
	//recursively call the two parts and find the min from them
	min := math.Min(divideAndConquer(pointsG[len(pointsG)/2:]), divideAndConquer(pointsG[:len(pointsG)/2]))
	if min == 0 {
		return 0
	}
	return math.Min(min, findYMin(pointsG, min))
}

func findMinDistY(points []Point) float64 {
	var min float64 = math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j-i <= 7 && j < len(points); j++ {
			min = math.Min(min, points[i].calculateDistance(points[j]))
		}
	}
	return min
}

func findYMin(pointsG []pointsGroup, min float64) float64 {

	//find the right and the left max x values
	diffL := pointsG[len(pointsG)/2].pointWithSameX[0].x - int(min)
	diffR := pointsG[len(pointsG)/2].pointWithSameX[0].x + int(min)

	//counters for right and left
	var counter1 int = len(pointsG)/2 - 1
	var counter2 int = len(pointsG)/2 + 1

	yCoordinates := make([]Point, 0, 100)

	for i := 0; i < len(pointsG[len(pointsG)/2].pointWithSameX); i++ {
		yCoordinates = append(yCoordinates, pointsG[len(pointsG)/2].pointWithSameX[i])
	}

	// fmt.Println(yCoordinates)

	for counter1 >= 0 && pointsG[counter1].pointWithSameX[0].x >= diffL {
		yCoordinates = mergeDifferent(yCoordinates, pointsG[counter1].pointWithSameX)
		counter1--
	}

	for counter2 <= len(pointsG)-1 && pointsG[counter2].pointWithSameX[0].x <= diffR {
		yCoordinates = mergeDifferent(yCoordinates, pointsG[counter2].pointWithSameX)
		counter2++
	}

	// fmt.Println(yCoordinates)
	return math.Min(min, findMinDistY(yCoordinates))
}

func main() {
	//declaring variables
	var noOfPoints int
	fmt.Scan(&noOfPoints)

	//making the points slice and it's mirror
	points := make([]Point, 0, noOfPoints)

	//inputting the data in both the slices
	points = inputPoints(points, noOfPoints)

	//mergeSorting the points
	mergeSort(points)
	// fmt.Println(points)

	//forming groups with same x
	pointsG := breakThemIntoSections(points)
	// fmt.Println(pointsG)

	//printing the result
	fmt.Printf("%.4f\n", divideAndConquer(pointsG))

}
