package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{35, -50, 42, 14, -5, 86, -21, -123, 34, 3233, 2}
	fmt.Println(maxSum(input))

}

func maxSum(numbers []int) int {
	finalSum := 0
	sum := 0
	for i := range numbers {
		if numbers[i] < 0 && math.Abs(float64(numbers[i])) > float64(sum) {
			finalSum = int(math.Max(float64(sum), float64(finalSum)))
			sum = 0
			continue
		}
		sum = sum + numbers[i]
		finalSum = int(math.Max(float64(sum), float64(finalSum)))
	}

	return finalSum
}
