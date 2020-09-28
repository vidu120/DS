package main

import (
	"fmt"
	"math"
	"strconv"
)

func minAndMax(expression string, max, min *[15][15]int, i, j int) (int, int) {
	//for keeping track of the min and the max values
	var valMin, valMax int

	valMin = math.MaxInt64
	valMax = math.MinInt64

	var a, b, c, d int

	for k := i; k < j; k++ {
		switch expression[2*k+1] {
		case '+':
			a = max[i][k] + max[k+1][j]
			b = min[i][k] + max[k+1][j]
			c = min[i][k] + min[k+1][j]
			d = max[i][k] + min[k+1][j]
		case '*':
			a = max[i][k] * max[k+1][j]
			b = min[i][k] * max[k+1][j]
			c = min[i][k] * min[k+1][j]
			d = max[i][k] * min[k+1][j]
		case '-':
			a = max[i][k] - max[k+1][j]
			b = min[i][k] - max[k+1][j]
			c = min[i][k] - min[k+1][j]
			d = max[i][k] - min[k+1][j]
		}
		valMin = int(math.Min(float64(a), float64(valMin)))
		valMin = int(math.Min(float64(b), float64(valMin)))
		valMin = int(math.Min(float64(c), float64(valMin)))
		valMin = int(math.Min(float64(d), float64(valMin)))

		valMax = int(math.Max(float64(a), float64(valMax)))
		valMax = int(math.Max(float64(b), float64(valMax)))
		valMax = int(math.Max(float64(c), float64(valMax)))
		valMax = int(math.Max(float64(d), float64(valMax)))
	}

	return valMin, valMax
}

func findMaxValu(expression string, max, min *[15][15]int) int {
	var (
		nNumbers int = len(expression)/2 + 1
	)

	var convToI = func(a string) int {
		val, _ := strconv.Atoi(a)
		return val
	}

	for i := 0; i < nNumbers; i++ {
		min[i][i] = convToI(string(expression[2*i]))
		max[i][i] = convToI(string(expression[2*i]))
	}

	//storing the max values
	var j int
	for s := 1; s < nNumbers; s++ {
		for i := 0; i < nNumbers-s; i++ {
			j = i + s
			min[i][j], max[i][j] = minAndMax(expression, max, min, i, j)
		}
	}

	// fmt.Println(min)
	// fmt.Println(max)

	return max[0][nNumbers-1]
}

func main() {
	//our arithmetic expression
	var expression string
	fmt.Scan(&expression)

	//for keeping track of all the min and the max expressions
	var max [15][15]int = [15][15]int{}
	var min [15][15]int = [15][15]int{}
	//we don't need to create separate arrays for numbers and symbols cause we know where they occur
	fmt.Println(findMaxValu(expression, &max, &min))
}
