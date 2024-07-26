package annonymousfunctions

import (
	"fmt"
)

type transformFn func(int) int

// example of long custom type
type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumber := []int{5, 1, 2}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformFunction(&numbers)
	transformFn2 := getTransformFunction(&moreNumber)

	transformedNumbers := transformNumber(&numbers, transformFn1)
	moreTransfomredNumbers := transformNumber(&numbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransfomredNumbers)
}

// returns function as a value
func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// returns function as value
func getTransformFunction(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
