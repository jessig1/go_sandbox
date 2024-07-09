package main

import (
	"fmt"
	"math"
)

func main() {
	// cannot change value
	const inflationRate = 2.5
	// initially infered by go to be an int, cast as a float both variables need to be of same type
	var investmentAmount, years float64 = 1000, 10
	// cannot convert variable type :=
	expectedReturnRate := 5.5
	// & is a pointer
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureRealValue)
}
