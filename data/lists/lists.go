package main

import "fmt"

// lists

func main() {
	prices := []float64{13.23, 72.18}
	fmt.Println(prices[0:1])
	prices[1] = 2.22

	prices = append(prices, 3.32, 9.82, 12.21)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{32.52, 55.55, 23.82}
	// ... adds a list of floats to existing list of floats
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

//dynamic arrays

// func main() {
// 	prices := []float64{13.23, 72.18}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 2.22

// 	prices = append(prices, 3.32)
// 	fmt.Println(prices)
// 	// updatedPrices := append(prices, 7.93)
// 	// fmt.Println("updated prices: ", updatedPrices)
// }

// static arrays
// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	// array with 4 prices
// 	prices := [4]float64{19.21, 3.24, 6.12, 23.45}
// 	fmt.Println(prices)
// 	// sets 3rd item in productNames
// 	productNames[2] = "A movie"
// 	fmt.Println(productNames[2])

// 	fmt.Println(prices[2])

// 	// slice removes first and last items
// 	featurePrices := prices[1:]
// 	// overrides existing 2nd item
// 	featurePrices[1] = 13.34
// 	highlightedPrices := featurePrices[:1]
// 	highlightedPrices = highlightedPrices[:3]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
