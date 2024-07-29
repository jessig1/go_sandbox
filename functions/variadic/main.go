package main

import (
	"fmt"
)

func main() {
	number := []int{3, 5, 8}
	sum := sumup(1, 10, 14)
	anotherSum := sumup(1, number...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, number ...int) int {
	sum := 0

	for _, val := range number {
		sum += val
	}

	return sum
}
