package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	// array with 4 prices
	prices := [4]float64{19.21, 3.24, 6.12, 23.45}
	fmt.Println(prices)
	// sets 3rd item in productNames
	productNames[2] = "A movie"
	fmt.Println(productNames[2])

	fmt.Println(prices[2])

	// slice removes first and last items
	featurePrices := prices[1:]
	// overrides existing 2nd item
	featurePrices[1] = 13.34
	highlightedPrices := featurePrices[:1]
	highlightedPrices = highlightedPrices[:3]

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

}
