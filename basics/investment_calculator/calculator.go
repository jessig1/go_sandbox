package main

import (
	"fmt"
	"math"
)

// cannot change value
const inflationRate = 2.5

func main() {

	// initially infered by go to be an int, cast as a float both variables need to be of same type
	var investmentAmount, years float64
	// cannot convert variable type :=
	expectedReturnRate := 5.5
	//fmt.Print("Investment amount: ")
	outputText("Investment Amount: ")
	// & is a pointer
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// assigns value of multiple variables based on return value
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("future real value: %v", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	// backticks ` allow for multiline formatting
	//fmt.Printf(`Future value: %.2f\nfuture real value: %v`, futureValue, futureRealValue)
	//fmt.Println("future real value: ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

// second paranetheses is the return values (in this case 2 float 64 values)
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}
